package user

import (
	"encoding/json"
	"time"

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
func (us *UserService) GetCurrentRate() (string, error) {
	request := CurrentCurrencyRateRequest{
		RequestID: "123",
		UserID:    1,
		Currency:  "USD",
	}
	requestJson, err := request.ToJSON()
	if err != nil {
		return "", err
	}
	msg, err := us.nc.Request("currency.rate", requestJson, 5*time.Second)
	if err != nil {
		return "", err
	}
	resp := CurrentCurrencyRateResponse{}
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return "", err
	}

	return resp.Rate, nil
}
