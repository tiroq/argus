package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type UserService struct {
	nc *nats.Conn
	// storage UserStorage
	// rateAPI RateAPI
}

type Option func(*UserService)

func WithBus(nc *nats.Conn) Option {
	return func(us *UserService) {
		us.nc = nc
	}
}

// func NewUserService(storage UserStorage, rateAPI RateAPI) *UserService {
func NewUserService(options ...Option) *UserService {
	srvc := &UserService{}
	for _, option := range options {
		option(srvc)
	}
	return srvc
}

// SubscribeUser subscribes a user to daily rate updates
// func (us *UserService) SubscribeUser(userID int) error {
// 	return us.storage.AddUserSubscription(userID)
// }

// GetCurrentRate fetches the current rate from the rate API
func (us *UserService) GetCurrentRate(userId int64, task string) (*CurrentCurrencyRateResponse, error) {
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

	msg, err := us.nc.RequestWithContext(ctx, "bog.currency.rate", requestJson)
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
