package user

import (
	"encoding/json"
	"fmt"
)

type CurrencyHistoryRequest struct {
	RequestID string `json:"request_id"`
	UserID    int    `json:"user_id"`
	Code      string `json:"code"`
}

func (r *CurrencyHistoryRequest) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return data, nil
}

func (r *CurrencyHistoryRequest) FromJSON(data []byte) error {
	err := json.Unmarshal(data, r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request: %w", err)
	}
	return nil
}
