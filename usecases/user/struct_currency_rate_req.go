package user

import (
	"encoding/json"
	"fmt"
)

type CurrentCurrencyRateRequest struct {
	RequestID string `json:"request_id"`
	UserID    int    `json:"user_id"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func (r *CurrentCurrencyRateRequest) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return data, nil
}

func (r *CurrentCurrencyRateRequest) FromJSON(data []byte) error {
	err := json.Unmarshal(data, r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request: %w", err)
	}
	return nil
}
