package keyboards

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//Создание инлайн кнопок
var (
  Button1 = tgbotapi.NewInlineKeyboardButtonData("Тестирование🤬", "Тестирование")
  Button2 = tgbotapi.NewInlineKeyboardButtonData("Из мира игр🎮", "Из мира игр")
  Button3 = tgbotapi.NewInlineKeyboardButtonData("Искуственный интеллект🦾", "Искусственный интеллект")
)

//Создание инлайн клавиатуры
var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
  tgbotapi.NewInlineKeyboardRow(Button1),
  tgbotapi.NewInlineKeyboardRow(Button2),
  tgbotapi.NewInlineKeyboardRow(Button3),
)   