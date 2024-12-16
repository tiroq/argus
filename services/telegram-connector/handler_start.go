package telegram

import "gopkg.in/telebot.v4"

func (tb *TelegramBot) handleStartCommand(c telebot.Context) error {
	tb.logger.Info("New user", "user_id", c.Sender().ID)
	return c.Send("Welcome! Use /subscribe to get daily updates and /rate to get the latest rate.")
}
