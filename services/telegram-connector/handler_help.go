package telegram

import (
	"gopkg.in/telebot.v4"
)

func (h *TelegramBot) handleHelpCommand(c telebot.Context) error {
	// Define a custom reply keyboard
	// menu := &telebot.ReplyMarkup{
	// 	ReplyKeyboard: [][]telebot.ReplyButton{
	// 		{
	// 			{
	// 				Text: "Hello",
	// 			},
	// 			{
	// 				Text: "Goodbye",
	// 			},
	// 		},
	// 	},
	// 	ResizeKeyboard: true,
	// }
	// menu = &telebot.ReplyMarkup{ResizeKeyboard: true}
	// selector = &telebot.ReplyMarkup{}

	// // Reply buttons.
	// btnHelp = menu.Text("ℹ Help")
	// btnSettings = menu.Text("⚙ Settings")
	return c.Send("Welcome! Use /subscribe to get daily updates and /rate to get the latest rate.")
}
