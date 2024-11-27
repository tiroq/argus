package scraper

import (
	"encoding/json"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/usecases/user"
)

func (s *Scraper) handleCurrencyRate(msg *nats.Msg) {
	s.logger.Info("Received currency rate request", "rate", string(msg.Data))
	request := user.CurrentCurrencyRateRequest{}
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		s.logger.Error("Failed to unmarshal request",
			slog.String("error", err.Error()))
		return
	}

	response := user.CurrentCurrencyRateResponse{
		RequestID: request.RequestID,
		Rate:      "1.2345",
	}
	data, err := json.Marshal(response)
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
