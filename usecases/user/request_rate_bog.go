package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tiroq/argus/internal/endpoints"
)

// GetCurrentRate fetches the current rate from the rate API
func (us *UserService) GetCurrentBOGRate(userId int64, task string) (*CurrentCurrencyRateResponse, error) {
	// Split the string into "From" and "To" currencies
	parts := strings.Split(task, "2")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format for task: %s", task)
	}

	// Convert "From" and "To" currencies to uppercase
	fromCurrency := strings.ToUpper(parts[0])
	toCurrency := strings.ToUpper(parts[1])
	request := CurrentCurrencyRateRequest{
		RequestID: uuid.New().String(),
		UserID:    int(userId),
		From:      fromCurrency,
		To:        toCurrency,
	}
	requestJson, err := request.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // should always be called, not discarded, to prevent context leak

	msg, err := us.nc.RequestWithContext(ctx, endpoints.SVC_BOG, requestJson)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	resp := &CurrentCurrencyRateResponse{}
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return resp, nil
}
