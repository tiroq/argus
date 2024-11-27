package main

import (
	"log"

	"github.com/tiroq/argus/internal/config"
	scraper "github.com/tiroq/argus/services/data-scraper"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	srvc, err := scraper.New(cfg)
	if err != nil {
		log.Fatal("Failed to create scraper:", err)
	}

	srvc.Start()
}
