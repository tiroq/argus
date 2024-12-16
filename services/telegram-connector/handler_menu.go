package telegram

import (
	"gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleMenuCommand(c telebot.Context) error {
	h.logger.Info("Menu clicked", "user_id", c.Sender().ID)
	return c.Send("Menu updated!", h.menu.Menu)
}
