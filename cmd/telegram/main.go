package main

import (
	"log"

	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/internal/telegram"
	"github.com/tiroq/argus/usecases/user"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize storage and rate API (mock implementations here)
	// storage := NewMockStorage()      // Mock storage for users
	// rateAPI := NewMockRateAPI()      // Mock rate API for fetching rates

	// Initialize the user service
	// userService := user.NewUserService(storage, rateAPI)
	userService := user.NewUserService()

	// Initialize and start the Telegram bot
	// bot, err := telegram.NewTelegramBot(cfg.TelegramToken, userService)
	bot, err := telegram.New(cfg, userService)
	if err != nil {
		log.Fatal("Failed to create Telegram bot:", err)
	}

	bot.Start()
}
