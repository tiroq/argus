package main

import (
	"log/slog"
	"os"

	"github.com/tiroq/argus/internal/config"
	kutsige "github.com/tiroq/argus/services/kursi-ge"
)

func main() {
	slog.Info("Starting Kurse.GE currency rate service")

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config:", slog.String("error", err.Error()))
		os.Exit(1)
	}

	srvc, err := kutsige.New(cfg)
	if err != nil {
		slog.Error("Failed to create BoG currency rate service :", slog.String("error", err.Error()))
		os.Exit(1)
	}

	srvc.Start()
}
