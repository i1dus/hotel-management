package send_all_message_usecase

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type EmployeeRepository interface {
	ListEmployees(ctx context.Context) ([]domain.Employee, error)
}

type SendAllMessageUseCase struct {
	bot          *tele.Bot
	employeeRepo EmployeeRepository
}

func NewSendAllMessageUseCase(bot *tele.Bot, employeeRepo EmployeeRepository) *SendAllMessageUseCase {
	return &SendAllMessageUseCase{bot: bot, employeeRepo: employeeRepo}
}

func (uc *SendAllMessageUseCase) SendAllMessage(c tele.Context) error {
	message := c.Message().Text
	if len(message) <= len(domain.CommandSendAllMessage)+1 { // длина команды + 1 символ пробела
		return c.Send(usecase.ErrorMessage(errors.New("отправлено пустое сообщение")))
	}

	employees, err := uc.employeeRepo.ListEmployees(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	messageToSend := fmt.Sprintf("Получено сообщение от @%s:\n%s",
		c.Sender().Username, message[len(domain.CommandSendAllMessage)+1:])

	// Отправить работникам и собрать юзернеймы получателей
	usernames := make([]string, 0, len(employees))
	for _, employee := range employees {
		if employee.UserID != 0 {
			_, err = uc.bot.Send(&tele.User{ID: employee.UserID}, messageToSend)
			if err != nil {
				continue // Ошибка = не отправилось, не записываем юзернейм
			}
			usernames = append(usernames, employee.Username)
		}
	}

	var responseMessage strings.Builder
	responseMessage.WriteString("Сообщение было отправлено следующим сотрудникам:")

	if len(usernames) == 0 {
		responseMessage.WriteString("\nСотрудники не найдены")
		return c.Send(responseMessage.String())
	}

	for i, username := range usernames {
		responseMessage.WriteString(fmt.Sprintf("\n%d. @%s", i+1, username))
	}
	return c.Send(responseMessage.String())
}
