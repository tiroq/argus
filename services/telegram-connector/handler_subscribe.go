package telegram

import "github.com/tucnak/telebot"

func (tb *TelegramBot) handleSubscribeCommand(m *telebot.Message) {
	tb.logger.Info("Subscribing user", "user_id", m.Sender.ID)
	// if err := h.userService.SubscribeUser(m.Sender.ID); err != nil {
	// 	h.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	// } else {
	// 	h.bot.Send(m.Sender, "You've been subscribed to daily rate updates!")
	// }
	tb.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	tb.logger.Error("Failed to subscribe user", "user_id", m.Sender.ID)
}
