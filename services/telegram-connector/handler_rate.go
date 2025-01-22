package telegram

import (
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleRateCommand(c telebot.Context) error {
	h.logger.Info("Rate requested", "user_id", c.Sender().ID, "data", c.Data())
	rate, err := h.userService.GetCurrentRate(c.Sender().ID, c.Data())
	strRate := strconv.FormatFloat(rate.Rate, 'f', -1, 64)
	strAmount := strconv.FormatFloat(rate.Amount, 'f', -1, 64)
	strRateSelf := strconv.FormatFloat(rate.RateSelf, 'f', -1, 64)
	strAmountSelf := strconv.FormatFloat(rate.AmountSelf, 'f', -1, 64)
	if err != nil {
		h.logger.Error("Failed to retrieve the rate", "error", err)
		return c.Send("Failed to retrieve the rate.")
	} else {
		return c.Send(
			fmt.Sprintf("Rate for %s --> %s\niBank:\t%s (1000 %s = %s %s)\nMBank:\t%s (1000 %s = %s %s)",
				rate.From, rate.To,
				strRate, rate.From, strAmount, rate.To,
				strRateSelf, rate.From, strAmountSelf, rate.To,
			),
		)
	}
}
