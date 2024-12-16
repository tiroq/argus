package scraper

import (
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/usecases/user"
)

func (s *Scraper) handleCurrencyRate(msg *nats.Msg) {
	s.logger.Info("Received currency rate request", "rate", string(msg.Data))
	request := user.CurrentCurrencyRateRequest{}
	if err := request.FromJSON(msg.Data); err != nil {
		s.logger.Error("Failed to unmarshal request",
			slog.String("error", err.Error()))
		return
	}

	response := user.CurrentCurrencyRateResponse{
		RequestID: request.RequestID,
		Rate:      1.2345,
	}
	data, err := response.ToJSON()
	if err != nil {
		s.logger.Error("Failed to marshal response",
			slog.String("error", err.Error()))
		return
	}

	if err := msg.Respond(data); err != nil {
		s.logger.Error("Failed to respond to request",
			slog.String("error", err.Error()))
	}
}
