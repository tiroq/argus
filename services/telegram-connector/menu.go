package telegram

import (
	"gopkg.in/telebot.v4"
)

type Menu struct {
	Menu          *telebot.ReplyMarkup
	BoG           *telebot.ReplyButton
	Help          *telebot.ReplyButton
	InlineBoGMenu *InlineBoGMenu
}

func NewMenu() *Menu {
	menu := &Menu{
		Menu: &telebot.ReplyMarkup{
			// InlineKeyboard: [][]telebot.InlineButton{
			// 	{
			// 		{
			// 			Unique: "boG",
			// 			Text:   "⚡ BoG",
			// 		},
			// 		{
			// 			Unique: "help",
			// 			Text:   "ℹ Help",
			// 		},
			// 	},
			// },

			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		},
		BoG: &telebot.ReplyButton{
			Text: "⚡ BoG",
		},
		Help: &telebot.ReplyButton{
			Text: "ℹ Help",
		},
		InlineBoGMenu: NewInlineBoGMenu(),
	}

	menu.Menu.ReplyKeyboard = [][]telebot.ReplyButton{
		{
			*menu.BoG,
			*menu.Help,
		},
	}

	return menu
}

type InlineBoGMenu struct {
	Menu *telebot.ReplyMarkup
}

func (i *InlineBoGMenu) AddButtonsHandling(tb *TelegramBot) {
	for _, row := range i.Menu.InlineKeyboard {
		for _, btn := range row {
			// tb.bot.Handle(&btn, func(c telebot.Context) error {
			// 	return c.Send(btn.Text)
			// })
			tb.bot.Handle(&btn, tb.handleRateCommand)
		}
	}
}

func NewInlineBoGMenu() *InlineBoGMenu {
	return &InlineBoGMenu{
		Menu: &telebot.ReplyMarkup{
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
			InlineKeyboard: [][]telebot.InlineButton{
				{
					{
						Unique: "usd2gel",
						Text:   "USD -> GEL",
						Data:   "usd2gel",
					},
					{
						Unique: "gel2usd",
						Text:   "GEL -> USD",
						Data:   "gel2usd",
					},
				},
				{
					{
						Unique: "eur2gel",
						Text:   "EUR -> GEL",
						Data:   "eur2gel",
					},
					{
						Unique: "gel2eur",
						Text:   "GEL -> EUR",
						Data:   "gel2eur",
					},
				},
				{
					{
						Unique: "gbp2gel",
						Text:   "GBP -> GEL",
						Data:   "gbp2gel",
					},
					{
						Unique: "gel2gbp",
						Text:   "GEL -> GBP",
						Data:   "gel2gbp",
					},
				},
			},
		},
	}
}
