package telegram

import (
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleRateCommand(c telebot.Context) error {
	h.logger.Info("Rate requested", "user_id", c.Sender().ID, "data", c.Data())

	// BoG rate
	bogRate, err := h.userService.GetCurrentBOGRate(c.Sender().ID, c.Data())
	strBOGRate := strconv.FormatFloat(bogRate.Rate, 'f', -1, 64)
	strBOGAmount := strconv.FormatFloat(bogRate.Amount, 'f', -1, 64)
	strBOGRateSelf := strconv.FormatFloat(bogRate.RateSelf, 'f', -1, 64)
	strBOGAmountSelf := strconv.FormatFloat(bogRate.AmountSelf, 'f', -1, 64)

	// KursiGe rate
	kursiRate, err := h.userService.GetCurrentKursiGERate(c.Sender().ID, c.Data())
	if err != nil {
		h.logger.Error("Failed to retrieve the rate", "error", err)
		return c.Send("Failed to retrieve the rate.")
	}
	h.logger.Info("kursiRate", "kursiRate", kursiRate)
	strKursiRate := strconv.FormatFloat(kursiRate.Rate, 'f', -1, 64)
	strKursiAmount := strconv.FormatFloat(kursiRate.Amount, 'f', -1, 64)

	if err != nil {
		h.logger.Error("Failed to retrieve the rate", "error", err)
		return c.Send("Failed to retrieve the rate.")
	} else {
		return c.Send(
			fmt.Sprintf("Rate for %s --> %s\nBoG iBank:\t%s (1000 %s = %s %s)\nBoG MBank:\t%s (1000 %s = %s %s)\nKursiGe:\t%s (1000 %s = %s %s)",
				bogRate.From, bogRate.To,
				strBOGRate, bogRate.From, strBOGAmount, bogRate.To,
				strBOGRateSelf, bogRate.From, strBOGAmountSelf, bogRate.To,
				strKursiRate, bogRate.From, strKursiAmount, bogRate.To,
			),
		)
	}
}
