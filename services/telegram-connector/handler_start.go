package telegram

import "github.com/tucnak/telebot"

func (tb *TelegramBot) handleStartCommand(m *telebot.Message) {
	tb.logger.Info("New user", "user_id", m.Sender.ID)
	tb.bot.Send(m.Sender, "Welcome! Use /subscribe to get daily updates and /rate to get the latest rate.")
}
