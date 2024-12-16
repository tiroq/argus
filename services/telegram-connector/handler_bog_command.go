package telegram

import (
	"gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleBoGCommand(c telebot.Context) error {
	return c.Send("BoG", h.menu.InlineBoGMenu.Menu)
}
