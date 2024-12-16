package bogcurrencyrate

import (
	"log"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
)

type RateResponse struct {
	Data struct {
		Rate   float64 `json:"rateSelf"`
		Amount float64 `json:"amountSelf"`
	} `json:"data"`
}

type BoGRatesProvider struct {
	logger *slog.Logger
	nc     *nats.Conn
}

type Option func(*BoGRatesProvider)

func WithLogger(logger *slog.Logger) Option {
	return func(s *BoGRatesProvider) {
		s.logger = logger
	}
}

func WithBus(nc *nats.Conn) Option {
	return func(s *BoGRatesProvider) {
		s.nc = nc
	}
}

func New(cfg *config.Config) (*BoGRatesProvider, error) {
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
		return nil, err
	}
	s := &BoGRatesProvider{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		nc:     nc,
	}
	// for _, option := range options {
	// 	option(s)
	// }
	return s, nil
}

func (s *BoGRatesProvider) Start() {
	s.nc.Subscribe("bog.currency.rate", s.handleCurrencyRate)
	select {} // Block the main goroutine
}
