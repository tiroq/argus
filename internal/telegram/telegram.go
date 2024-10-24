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

func New(cfg *config.Config, userService *user.UserService) (*TelegramBot, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		logger.Error("Failed to connect to NATS server",
			slog.String("url", cfg.NatsUrl))
		return nil, err
	}
	defer nc.Close()

	logger.Info("Connected to NATS server")
	b, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		logger:      logger,
		bus:         nc,
		bot:         b,
		userService: userService,
	}, nil
}

func (tb *TelegramBot) Start() {
	tb.bot.Handle("/start", func(m *telebot.Message) {
		tb.logger.Info("New user", "user_id", m.Sender.ID)
		tb.bot.Send(m.Sender, "Welcome! Use /subscribe to get daily updates and /rate to get the latest rate.")
	})

	tb.bot.Handle("/subscribe", func(m *telebot.Message) {
		// if err := tb.userService.SubscribeUser(m.Sender.ID); err != nil {
		// 	tb.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
		// } else {
		// 	tb.bot.Send(m.Sender, "You've been subscribed to daily rate updates!")
		// }
		tb.bot.Send(m.Sender, "Failed to subscribe you. Try again later.")
	})

	tb.bot.Handle("/rate", func(m *telebot.Message) {
		rate, err := tb.userService.GetCurrentRate()
		if err != nil {
			tb.bot.Send(m.Sender, "Failed to retrieve the rate.")
		} else {
			tb.bot.Send(m.Sender, "The current rate is: "+rate)
		}
		// tb.bot.Send(m.Sender, "Failed to retrieve the rate.")
	})

	tb.bot.Start()
}
