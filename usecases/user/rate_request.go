package user

import (
	"encoding/json"
	"fmt"
)

type CurrentCurrencyRateRequest struct {
	RequestID string `json:"request_id"`
	UserID    int    `json:"user_id"`
	Currency  string `json:"currency"`
}

func (r CurrentCurrencyRateRequest) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return data, nil
}

type CurrentCurrencyRateResponse struct {
	RequestID string `json:"request_id"`
	Rate      string `json:"rate"`
}
