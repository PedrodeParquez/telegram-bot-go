package keyboards

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Создание клавиатуры с обычной
var ButtonsKeyboard = []tgbotapi.KeyboardButton{
  tgbotapi.NewKeyboardButton("Кнопка 1"),
  tgbotapi.NewKeyboardButton("Кнопка 2"),
  tgbotapi.NewKeyboardButton("Кнопка 3"),
}

var Keyboard = tgbotapi.NewReplyKeyboard(ButtonsKeyboard)
