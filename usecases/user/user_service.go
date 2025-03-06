package user

import (
	"log/slog"

	"github.com/nats-io/nats.go"
)

type UserService struct {
	nc     *nats.Conn
	logger *slog.Logger
	// storage UserStorage
	// rateAPI RateAPI
}

type Option func(*UserService)

func WithBus(nc *nats.Conn) Option {
	return func(us *UserService) {
		us.nc = nc
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(us *UserService) {
		us.logger = logger
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
