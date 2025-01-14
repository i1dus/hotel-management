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
	ErrRoomsNotFound = errors.New("номера не найдены")
	ErrRoomNotFound  = errors.New("номер не найден")
)

type RoomRepository struct {
	conn *pgx.Conn
}

func NewRoomRepository(conn *pgx.Conn) *RoomRepository {
	return &RoomRepository{conn: conn}
}

func (r *RoomRepository) AddRoom(ctx context.Context, room domain.Room) error {
	modelRoom := model.Rooms{
		Number: room.Number,
		Type:   int32(room.Type),
		Price:  int32(room.Price),
	}

	stmt, args := table.Rooms.
		INSERT(table.Rooms.AllColumns.Except(table.Rooms.ID)).
		MODEL(modelRoom).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}

func (r *RoomRepository) ListRooms(ctx context.Context) ([]domain.Room, error) {
	stmt, args := postgres.SELECT(
		table.Rooms.AllColumns).
		FROM(table.Rooms).
		ORDER_BY(table.Rooms.Number).
		Sql()

	var modelRooms []model.Rooms

	rows, err := r.conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoomsNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var room model.Rooms
		if err := rows.Scan(&room.ID, &room.Number, &room.Type, &room.Price, &room.Cleaned, &room.Description); err != nil {
			return nil, err
		}
		modelRooms = append(modelRooms, room)
	}

	rooms := lo.Map(modelRooms, func(modelRoom model.Rooms, index int) domain.Room {
		return domain.Room{
			Number:      modelRoom.Number,
			Type:        domain.RoomCategory(modelRoom.Type),
			Price:       int(modelRoom.Price),
			Cleaned:     modelRoom.Cleaned,
			Description: lo.FromPtr(modelRoom.Description),
		}
	})

	return rooms, nil
}

func (r *RoomRepository) ChangeRoomPrice(ctx context.Context, number string, price int) error {
	stmt, args := table.Rooms.
		UPDATE(table.Rooms.Price).
		SET(postgres.Int(int64(price))).
		WHERE(table.Rooms.Number.EQ(postgres.String(number))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrRoomNotFound
	}
	return nil
}

func (r *RoomRepository) IsRoomExist(ctx context.Context, number string) (bool, error) {
	stmt, args := postgres.
		SELECT(postgres.COUNT(postgres.STAR)).
		FROM(table.Rooms).
		WHERE(table.Rooms.Number.EQ(postgres.String(number))).Sql()

	var count int64
	err := r.conn.QueryRow(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (r *RoomRepository) ChangeRoomCleaned(ctx context.Context, number string, cleaned bool) error {
	stmt, args := table.Rooms.
		UPDATE(table.Rooms.Cleaned).
		SET(postgres.Bool(cleaned)).
		WHERE(table.Rooms.Number.EQ(postgres.String(number))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrRoomNotFound
	}
	return nil
}

func (r *RoomRepository) ChangeRoomDescription(ctx context.Context, number string, description string) error {
	stmt, args := table.Rooms.
		UPDATE(table.Rooms.Description).
		SET(postgres.String(description)).
		WHERE(table.Rooms.Number.EQ(postgres.String(number))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrRoomNotFound
	}
	return nil
}
