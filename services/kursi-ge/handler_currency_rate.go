package kursige

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/usecases/user"
)

func (s *KursiGERatesProvider) handleCurrencyRate(msg *nats.Msg) {
	s.logger.Info("Received currency rate request", "rate", string(msg.Data))
	var request user.CurrentCurrencyRateRequest
	if err := request.FromJSON(msg.Data); err != nil {
		s.logger.Error("Failed to unmarshal request",
			slog.String("error", err.Error()))
		return
	}

	s.logger.Info("Getting currency rate", "from", request.From, "to", request.To)

	rate, err := s.getCurrencyRate()
	if err != nil {
		s.logger.Error("Failed to get currency rate",
			slog.String("error", err.Error()))
		return
	}

	s.logger.Info("Sending currency rate response", "rate", rate)

	response := user.CurrentCurrencyRateResponse{}
	for _, rate := range rate.Data {
		// i.e. GEL --> USD
		if rate.SecondaryCurrencyCode == request.To && rate.BaseCurrencyCode == request.From {
			response = user.CurrentCurrencyRateResponse{
				RequestID:  request.RequestID,
				Amount:     1000 / rate.SellRate,
				Rate:       rate.SellRate,
				AmountSelf: 0,
				RateSelf:   0,
				From:       request.From,
				To:         request.To,
			}
			break
		}

		// i.e. USD --> GEL
		if rate.SecondaryCurrencyCode == request.From && rate.BaseCurrencyCode == request.To {
			response = user.CurrentCurrencyRateResponse{
				RequestID:  request.RequestID,
				Amount:     1000 * rate.BuyRate,
				Rate:       rate.BuyRate,
				AmountSelf: 0,
				RateSelf:   0,
				From:       request.From,
				To:         request.To,
			}
			break
		}
	}
	// RequestID:  request.RequestID,
	// Amount:     rate.Data.Amount,
	// Rate:       rate.Data.Rate,
	// AmountSelf: rate.Data.AmountSelf,
	// RateSelf:   rate.Data.RateSelf,
	// From:       request.From,
	// To:         request.To}
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

func (s *KursiGERatesProvider) getCurrencyRate() (*RateResponse, error) {

	// Check cache
	mutex.Lock()
	if item, found := cache["data"]; found && time.Since(item.Timestamp) < cacheDuration {
		mutex.Unlock()
		s.logger.Info("Using cached rate", "rate", item.RateInfo)
		return &item.RateInfo, nil
	}
	mutex.Unlock()

	// Make request to BoG API
	url := "https://api.kursi.ge/api/public/currencies"
	s.logger.Info("Making request to BoG API", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Received currency rate response", "rate", string(body))

	var rateResponse RateResponse
	err = json.Unmarshal(body, &rateResponse.Data)
	if err != nil {
		return nil, err
	}

	// Update cache
	mutex.Lock()
	cache["data"] = CacheItem{
		RateInfo:  rateResponse,
		Timestamp: time.Now(),
	}
	mutex.Unlock()

	s.logger.Info("Response rate", "rate", rateResponse)
	return &rateResponse, nil
}
