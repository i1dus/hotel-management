package employee_handler

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type AddEmployeeUseCase interface {
	AddEmployee(c tele.Context) error
}

type RemoveEmployeeUseCase interface {
	RemoveEmployee(c tele.Context) error
}

type ListEmployeesUseCase interface {
	ListEmployees(c tele.Context) error
}

type SendAllMessageUseCase interface {
	SendAllMessage(c tele.Context) error
}

// EmployeeHandler has employees managing commands
type EmployeeHandler struct {
	bot                   *tele.Bot
	addEmployeeUseCase    AddEmployeeUseCase
	removeEmployeeUseCase RemoveEmployeeUseCase
	listEmployeesUseCase  ListEmployeesUseCase
	sendAllMessageUseCase SendAllMessageUseCase
}

func NewEmployeeHandler(bot *tele.Bot, addEmployeeUseCase AddEmployeeUseCase, removeEmployeeUseCase RemoveEmployeeUseCase, listEmployeesUseCase ListEmployeesUseCase, sendAllMessageUseCase SendAllMessageUseCase) *EmployeeHandler {
	return &EmployeeHandler{bot: bot, addEmployeeUseCase: addEmployeeUseCase, removeEmployeeUseCase: removeEmployeeUseCase, listEmployeesUseCase: listEmployeesUseCase, sendAllMessageUseCase: sendAllMessageUseCase}
}

func (h *EmployeeHandler) RegisterHandlers() {
	h.bot.Handle(domain.CommandAddEmployee, h.addEmployeeUseCase.AddEmployee)
	h.bot.Handle(domain.CommandRemoveEmployee, h.removeEmployeeUseCase.RemoveEmployee)
	h.bot.Handle(domain.CommandListEmployees, h.listEmployeesUseCase.ListEmployees)
	h.bot.Handle(domain.CommandSendAllMessage, h.sendAllMessageUseCase.SendAllMessage)
}
