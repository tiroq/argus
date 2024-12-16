package base

import (
	"encoding/json"
	"fmt"
)

type Message struct{}

func (m *Message) ToJSON() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return data, nil
}

func (m *Message) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request: %w", err)
	}
	return nil
}
