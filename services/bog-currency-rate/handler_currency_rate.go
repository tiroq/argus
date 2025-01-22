package bogcurrencyrate

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/usecases/user"
)

func (s *BoGRatesProvider) handleCurrencyRate(msg *nats.Msg) {
	s.logger.Info("Received currency rate request", "rate", string(msg.Data))
	var request user.CurrentCurrencyRateRequest
	if err := request.FromJSON(msg.Data); err != nil {
		s.logger.Error("Failed to unmarshal request",
			slog.String("error", err.Error()))
		return
	}

	s.logger.Info("Getting currency rate", "from", request.From, "to", request.To)

	rate, err := s.getCurrencyRate(request)
	if err != nil {
		s.logger.Error("Failed to get currency rate",
			slog.String("error", err.Error()))
		return
	}

	s.logger.Info("Sending currency rate response", "rate", rate)

	response := user.CurrentCurrencyRateResponse{
		RequestID:  request.RequestID,
		Amount:     rate.Data.Amount,
		Rate:       rate.Data.Rate,
		AmountSelf: rate.Data.AmountSelf,
		RateSelf:   rate.Data.RateSelf,
		From:       request.From,
		To:         request.To}
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

func (s *BoGRatesProvider) getCurrencyRate(request user.CurrentCurrencyRateRequest) (*RateResponse, error) {
	cacheKey := fmt.Sprintf("%s_%s", request.From, request.To)

	// Check cache
	mutex.Lock()
	if item, found := cache[cacheKey]; found && time.Since(item.Timestamp) < cacheDuration {
		mutex.Unlock()
		s.logger.Info("Using cached rate", "rate", item.RateInfo)
		return &item.RateInfo, nil
	}
	mutex.Unlock()

	// Make request to BoG API
	url := fmt.Sprintf("https://bankofgeorgia.ge/api/currencies/convert/%s/%s?amountFrom=1000&amountTo", request.From, request.To)
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
	err = json.Unmarshal(body, &rateResponse)
	if err != nil {
		return nil, err
	}

	// Update cache
	mutex.Lock()
	cache[cacheKey] = CacheItem{
		RateInfo:  rateResponse,
		Timestamp: time.Now(),
	}
	mutex.Unlock()

	s.logger.Info("Response rate", "rate", rateResponse)
	return &rateResponse, nil
}
