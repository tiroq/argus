package nbg

import (
	"log"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/internal/endpoints"
)

type NBGHistoryProvider struct {
	logger *slog.Logger
	nc     *nats.Conn
}

type Option func(*NBGHistoryProvider)

func WithBus(nc *nats.Conn) Option {
	return func(s *NBGHistoryProvider) {
		s.nc = nc
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(s *NBGHistoryProvider) {
		s.logger = logger
	}
}

func New(cfg *config.Config) (*NBGHistoryProvider, error) {
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
		return nil, err
	}
	s := &NBGHistoryProvider{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		nc:     nc,
	}
	// for _, option := range options {
	// 	option(s)
	// }
	return s, nil
}

func (s *NBGHistoryProvider) Start() {
	s.logger.Info("Starting KursiGE service")
	s.nc.Subscribe(endpoints.SVC_NBG_HIST, s.handleCurrencyHistory)
	s.logger.Info("KursiGE service started")
	select {} // Block the main goroutine

}
