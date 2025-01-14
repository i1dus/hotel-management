package client_handler

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type AddClientUseCase interface {
	AddClient(c tele.Context) error
}

// ClientHandler has client managing commands
type ClientHandler struct {
	bot              *tele.Bot
	addClientUseCase AddClientUseCase
}

func NewClientHandler(bot *tele.Bot, addClientUseCase AddClientUseCase) *ClientHandler {
	return &ClientHandler{bot: bot, addClientUseCase: addClientUseCase}
}

func (h *ClientHandler) RegisterHandlers() {
	h.bot.Handle(domain.CommandAddClient, h.addClientUseCase.AddClient)
}
