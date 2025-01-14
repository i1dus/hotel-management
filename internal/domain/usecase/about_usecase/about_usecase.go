package about_usecase

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type AboutUseCase struct{}

var (
	selector  = &tele.ReplyMarkup{}
	btnReload = selector.URL("Ссылка на GitHub проекта", domain.AboutLink)
)

func NewAboutUseCase() *AboutUseCase {
	selector.Inline(
		selector.Row(btnReload),
	)
	return &AboutUseCase{}
}

func (uc *AboutUseCase) About(c tele.Context) error {
	return c.Send(domain.AboutMsg, selector)
}
