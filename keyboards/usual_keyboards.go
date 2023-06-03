package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Создание клавиатуры с обычной
var ButtonsKeyboard = []tgbotapi.KeyboardButton {
	tgbotapi.NewKeyboardButton("Удалить клавиатуру"),
	tgbotapi.NewKeyboardButton("Восстановить клавиатуру"),
}

var UsualKeyboard = tgbotapi.NewReplyKeyboard(ButtonsKeyboard)
