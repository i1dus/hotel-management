package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"hotel-management/internal/domain"
	"hotel-management/internal/gen/hotel_management/public/model"
	"hotel-management/internal/gen/hotel_management/public/table"
)

var (
	ErrEmployeeNotFound  = errors.New("сотрудник не найден")
	ErrEmployeesNotFound = errors.New("сотрудники не найдены")
)

type EmployeeRepository struct {
	conn *pgx.Conn
}

func NewEmployeeRepository(conn *pgx.Conn) *EmployeeRepository {
	return &EmployeeRepository{conn: conn}
}

func (r *EmployeeRepository) AddEmployee(ctx context.Context, employee domain.Employee) error {
	modelEmployee := model.Employees{
		Username: employee.Username,
		Name:     &employee.Name,
		Position: int32(employee.Position),
	}

	stmt, args := table.Employees.
		INSERT(table.Employees.AllColumns.Except(table.Employees.ID)).
		MODEL(modelEmployee).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}

func (r *EmployeeRepository) RemoveEmployee(ctx context.Context, username string) error {
	stmt, args := table.Employees.
		DELETE().
		WHERE(table.Employees.Username.EQ(postgres.String(username))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}
	return nil
}

func (r *EmployeeRepository) ListEmployees(ctx context.Context) ([]domain.Employee, error) {
	stmt, args := postgres.SELECT(
		table.Employees.AllColumns).
		FROM(table.Employees).Sql()

	var modelEmployees []model.Employees

	rows, err := r.conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployeesNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee model.Employees
		if err := rows.Scan(&employee.ID, &employee.Username, &employee.Name, &employee.Position, &employee.UserID); err != nil {
			return nil, err
		}
		modelEmployees = append(modelEmployees, employee)
	}

	employees := lo.Map(modelEmployees, func(modelEmployee model.Employees, index int) domain.Employee {
		return domain.Employee{
			Username: modelEmployee.Username,
			Name:     lo.FromPtr(modelEmployee.Name),
			Position: domain.Position(modelEmployee.Position),
			UserID: func(userID *int32) int64 {
				if userID == nil {
					return 0
				}
				return int64(*userID)
			}(modelEmployee.UserID),
		}
	})

	return employees, nil
}

func (r *EmployeeRepository) IsEmployeeWithPositions(ctx context.Context, username string, positions []domain.Position) (bool, error) {
	positionIDs := lo.Map(positions, func(item domain.Position, index int) postgres.Expression {
		return postgres.Int(int64(item))
	})

	stmt, args := postgres.SELECT(
		postgres.COUNT(postgres.STAR)).
		FROM(table.Employees).
		WHERE(table.Employees.Username.EQ(postgres.String(username)).
			AND(table.Employees.Position.IN(positionIDs...))).
		Sql()

	var count int64
	err := r.conn.QueryRow(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func (r *EmployeeRepository) UpsertEmployeeUserID(ctx context.Context, username string, userID int) error {
	stmt, args := table.Employees.
		UPDATE(table.Employees.UserID).
		SET(postgres.Int(int64(userID))).
		WHERE(table.Employees.Username.EQ(postgres.String(username))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}
	return nil
}
