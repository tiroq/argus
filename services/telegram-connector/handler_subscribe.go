package telegram

import "gopkg.in/telebot.v4"

func (tb *TelegramBot) handleSubscribeCommand(c telebot.Context) error {
	tb.logger.Info("Subscribing user", "user_id", c.Sender().ID)
	// if err := h.userService.SubscribeUser(c.Sender().ID); err != nil {
	// 	h.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	// } else {
	// 	h.bot.Send(m.Sender, "You've been subscribed to daily rate updates!")
	// }
	c.Send("Failed to subscribe you. Try again later.")
	tb.logger.Error("Failed to subscribe user", "user_id", c.Sender().ID)
	return nil
}
