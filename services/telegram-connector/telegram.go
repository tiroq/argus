package telegram

import (
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tiroq/argus/internal/config"
	"github.com/tiroq/argus/usecases/user"
	"gopkg.in/telebot.v4"
)

type TelegramBot struct {
	logger      *slog.Logger
	bus         *nats.Conn
	bot         *telebot.Bot
	userService *user.UserService
	menu        *Menu
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
		menu:        NewMenu(),
	}, nil
}

func (tb *TelegramBot) Start() {
	tb.bot.Handle("/start", tb.handleStartCommand)
	tb.bot.Handle("/menu", tb.handleMenuCommand)
	tb.bot.Handle("/help", tb.handleHelpCommand)
	tb.bot.Handle("/subscribe", tb.handleSubscribeCommand)
	tb.bot.Handle("/unsubscribe", tb.handleUnSubscribeCommand)
	tb.bot.Handle("/rate", tb.handleRateCommand)

	// Buttons
	tb.bot.Handle(tb.menu.BoG, tb.handleBoGCommand)
	tb.bot.Handle(tb.menu.Help, tb.handleHelpCommand)
	tb.menu.InlineBoGMenu.AddButtonsHandling(tb)
	// tb.bot.Handle(&inlineMenu, tb.handleInlineMenu)
	tb.bot.Start()
}
