package user

import (
	"encoding/json"
	"fmt"

	"github.com/tiroq/argus/services/nbg"
)

type CurrencyHistoryResponse struct {
	RequestID string `json:"request_id"`
	Data      []nbg.HistEntry
}

func (r *CurrencyHistoryResponse) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return data, nil
}

func (r *CurrencyHistoryResponse) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}
