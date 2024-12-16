package main

import (
	"log/slog"
	"os"

	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/services/telegram-connector"
)

func main() {
	slog.Info("Starting Telegram bot connector")
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config:", slog.String("error", err.Error()))
		os.Exit(1)
	}

	bot, err := telegram.New(cfg)
	if err != nil {
		slog.Error("Failed to create Telegram bot:", slog.String("error", err.Error()))
		os.Exit(1)
	}

	bot.Start()
}
