package scraper

import (
	"log"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
)

type Scraper struct {
	logger *slog.Logger
	nc     *nats.Conn
}

type Option func(*Scraper)

func WithLogger(logger *slog.Logger) Option {
	return func(s *Scraper) {
		s.logger = logger
	}
}

func WithBus(nc *nats.Conn) Option {
	return func(s *Scraper) {
		s.nc = nc
	}
}

func New(cfg *config.Config) (*Scraper, error) {
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
		return nil, err
	}
	s := &Scraper{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		nc:     nc,
	}
	// for _, option := range options {
	// 	option(s)
	// }
	return s, nil
}

func (s *Scraper) Start() {
	s.nc.Subscribe("currency.rate", s.handleCurrencyRate)
	select {} // Block the main goroutine
}
