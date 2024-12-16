package telegram

import (
	telebot "gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleInlineMenu(c telebot.Context) {
	h.logger.Info("Inline menu clicked", "user_id", c.Sender().ID)
}
