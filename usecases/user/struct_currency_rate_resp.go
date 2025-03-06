package user

import (
	"encoding/json"
	"fmt"
)

type CurrentCurrencyRateResponse struct {
	RequestID  string  `json:"request_id"`
	Rate       float64 `json:"rate"`
	Amount     float64 `json:"amount"`
	RateSelf   float64 `json:"rateSelf"`
	AmountSelf float64 `json:"amountSelf"`
	From       string  `json:"from"`
	To         string  `json:"to"`
}

func (r CurrentCurrencyRateResponse) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return data, nil
}

func (r CurrentCurrencyRateResponse) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}
