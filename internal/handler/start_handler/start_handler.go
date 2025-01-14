package start_handler

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type HelpUseCase interface {
	Help(c tele.Context) error
}

type StartUseCase interface {
	Start(c tele.Context) error
}

type AboutUseCase interface {
	About(c tele.Context) error
}

// StartHandler is responsible to establish connection between bot & user
type StartHandler struct {
	bot          *tele.Bot
	menuWrapper  *MenuWrapper
	helpUseCase  HelpUseCase
	aboutUseCase AboutUseCase
	startUseCase StartUseCase
}

func NewStartHandler(bot *tele.Bot, menuWrapper *MenuWrapper, helpUseCase HelpUseCase, aboutUseCase AboutUseCase, startUseCase StartUseCase) *StartHandler {
	return &StartHandler{bot: bot, menuWrapper: menuWrapper, helpUseCase: helpUseCase, aboutUseCase: aboutUseCase, startUseCase: startUseCase}
}

func (h *StartHandler) RegisterHandlers() {
	h.bot.Handle(domain.CommandStart, h.startUseCase.Start)
	h.bot.Handle(h.menuWrapper.BtnHelp, h.helpUseCase.Help)
	h.bot.Handle(h.menuWrapper.BtnAbout, h.aboutUseCase.About)
}
