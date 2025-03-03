package kursige

import (
	"log"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/internal/endpoints"
)

type BankRates struct {
	TBC struct {
		BuyRate  float64 `json:"buyRate"`
		SellRate float64 `json:"sellRate"`
	} `json:"TBC"`
	BOG struct {
		BuyRate  float64 `json:"buyRate"`
		SellRate float64 `json:"sellRate"`
	} `json:"BOG"`
}

type RateInfo struct {
	BaseCurrencyCode      string    `json:"baseCurrencyCode"`
	SecondaryCurrencyCode string    `json:"secondaryCurrencyCode"`
	Name                  string    `json:"name"`
	NbgRate               float64   `json:"nbgRate"`
	Diff                  float64   `json:"diff"`
	BuyRate               float64   `json:"buyRate"`
	SellRate              float64   `json:"sellRate"`
	BankRates             BankRates `json:"bankRates"`
}

type RateResponse struct {
	Data []RateInfo
}

type KursiGERatesProvider struct {
	logger *slog.Logger
	nc     *nats.Conn
}

type Option func(*KursiGERatesProvider)

func WithLogger(logger *slog.Logger) Option {
	return func(s *KursiGERatesProvider) {
		s.logger = logger
	}
}

func WithBus(nc *nats.Conn) Option {
	return func(s *KursiGERatesProvider) {
		s.nc = nc
	}
}

func New(cfg *config.Config) (*KursiGERatesProvider, error) {
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
		return nil, err
	}
	s := &KursiGERatesProvider{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		nc:     nc,
	}
	// for _, option := range options {
	// 	option(s)
	// }
	return s, nil
}

func (s *KursiGERatesProvider) Start() {
	s.logger.Info("Starting KursiGE service")
	s.nc.Subscribe(endpoints.SEVICE_KURSI_GE, s.handleCurrencyRate)
	s.logger.Info("KursiGE service started")
	select {} // Block the main goroutine
}
