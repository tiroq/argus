package telegram

import "github.com/tucnak/telebot"

func (h *TelegramBot) handleUnSubscribeCommand(m *telebot.Message) {
	// if err := h.userService.SubscribeUser(m.Sender.ID); err != nil {
	// 	h.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	// } else {
	// 	h.bot.Send(m.Sender, "You've been subscribed to daily rate updates!")
	// }
	h.bot.Send(m.Sender, "Failed to unsubscribe you. Try again later.")
}
