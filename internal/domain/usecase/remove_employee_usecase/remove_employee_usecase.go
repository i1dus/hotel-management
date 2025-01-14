package remove_employee_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type EmployeeRepository interface {
	RemoveEmployee(ctx context.Context, username string) error
}

type RemoveEmployeeUseCase struct {
	employeeRepo EmployeeRepository
}

func NewRemoveEmployeeUseCase(employeeRepo EmployeeRepository) *RemoveEmployeeUseCase {
	return &RemoveEmployeeUseCase{employeeRepo: employeeRepo}
}

func (uc *RemoveEmployeeUseCase) RemoveEmployee(c tele.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Должен быть 1 аргумент: TG-логин")
	}

	// ТГ-логин
	username := args[0]
	if len(username) > 1 && !strings.HasPrefix(username, "@") {
		return c.Send("Логин должен начинаться с '@'")
	}

	err := uc.employeeRepo.RemoveEmployee(context.Background(), username[1:])
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Сотрудник успешно удален!")
}
