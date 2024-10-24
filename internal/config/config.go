package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token   string
	Admin   string
	NatsUrl string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN environment variable not set")
	}

	admin := os.Getenv("TELEGRAM_ADMIN")
	if admin == "" {
		return nil, fmt.Errorf("TELEGRAM_ADMIN environment variable not set")
	}

	nats_url := os.Getenv("NATS_URL")
	if nats_url == "" {
		return nil, fmt.Errorf("NATS_URL environment variable not set")
	}
	return &Config{
		Token:   token,
		Admin:   admin,
		NatsUrl: nats_url,
	}, nil
}
