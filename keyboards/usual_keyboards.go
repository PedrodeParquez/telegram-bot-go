package keyboards

import (
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ButtonsKeyboard = []tgbotapi.KeyboardButton{
	tgbotapi.NewKeyboardButton("Кнопка 1"),
	tgbotapi.NewKeyboardButton("Кнопка 2"),
}

var UsualKeyboard = tgbotapi.NewReplyKeyboard(ButtonsKeyboard)

func ButtonOneClick(bot *tgbotapi.BotAPI, userID int64) {
	message.SendStringMessage(bot, userID, "Вы нажали кнопку 1")
	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &MainInlineKeyboard)
}

func ButtonTwoClick(bot *tgbotapi.BotAPI, userID int64) {
	message.SendStringMessage(bot, userID, "Вы нажали кнопку 2")
	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &MainInlineKeyboard)
}
