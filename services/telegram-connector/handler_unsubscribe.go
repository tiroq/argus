package telegram

import "gopkg.in/telebot.v4"

func (h *TelegramBot) handleUnSubscribeCommand(c telebot.Context) error {
	// if err := h.userService.SubscribeUser(c.Sender().ID); err != nil {
	// 	h.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	// } else {
	// 	h.bot.Send(m.Sender, "You've been subscribed to daily rate updates!")
	// }
	return c.Send("Failed to unsubscribe you. Try again later.")
}
