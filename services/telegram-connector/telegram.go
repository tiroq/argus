package telegram

import (
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/usecases/user"
	"github.com/tucnak/telebot"
)

type TelegramBot struct {
	logger      *slog.Logger
	bus         *nats.Conn
	bot         *telebot.Bot
	userService *user.UserService
}

func New(cfg *config.Config) (*TelegramBot, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		logger.Error("Failed to connect to NATS server",
			slog.String("url", cfg.NatsUrl))
		return nil, err
	}

	logger.Info("Connected to NATS server",
		slog.String("url", cfg.NatsUrl))
	b, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		logger.Error("Failed to create Telegram bot",
			slog.String("error", err.Error()))
		return nil, err
	}

	userService := user.NewUserService(user.WithBus(nc))
	// if err != nil {
	// 	logger.Error("Failed to create user service",
	// 		slog.String("error", err.Error()))
	// 	return nil, err
	// }
	return &TelegramBot{
		logger:      logger,
		bus:         nc,
		bot:         b,
		userService: userService,
	}, nil
}

func (tb *TelegramBot) Start() {
	tb.bot.Handle("/start", tb.handleStartCommand)
	tb.bot.Handle("/subscribe", tb.handleSubscribeCommand)
	tb.bot.Handle("/unsubscribe", tb.handleUnSubscribeCommand)
	tb.bot.Handle("/rate", tb.handleRateCommand)
	tb.bot.Start()
}
