package main

import (
	"log/slog"
	"os"

	"github.com/tiroq/argus/internal/config"
	bogcurrencyrate "github.com/tiroq/argus/services/bog-currency-rate"
)

func main() {
	slog.Info("Starting BoG currency rate service")

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config:", slog.String("error", err.Error()))
		os.Exit(1)
	}

	srvc, err := bogcurrencyrate.New(cfg)
	if err != nil {
		slog.Error("Failed to create BoG currency rate service :", slog.String("error", err.Error()))
		os.Exit(1)
	}

	srvc.Start()
}
