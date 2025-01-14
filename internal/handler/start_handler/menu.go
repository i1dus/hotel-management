package start_handler

import tele "gopkg.in/telebot.v4"

type MenuWrapper struct {
	Menu     *tele.ReplyMarkup
	BtnHelp  *tele.Btn
	BtnAbout *tele.Btn
}

func NewMenuWrapper() *MenuWrapper {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnHelp := menu.Text("ℹ️ Помощь по командам")
	btnAbout := menu.Text("⚡ Подробнее о сервисе")
	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnAbout),
	)

	return &MenuWrapper{
		Menu:     menu,
		BtnHelp:  &btnHelp,
		BtnAbout: &btnAbout,
	}
}
