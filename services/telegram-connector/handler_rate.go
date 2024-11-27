package telegram

import "github.com/tucnak/telebot"

func (h *TelegramBot) handleRateCommand(m *telebot.Message) {
	rate, err := h.userService.GetCurrentRate()
	if err != nil {
		h.logger.Error("Failed to retrieve the rate", "error", err)
		h.bot.Send(m.Sender, "Failed to retrieve the rate.")
	} else {
		h.bot.Send(m.Sender, "The current rate is: "+rate)
	}
}
