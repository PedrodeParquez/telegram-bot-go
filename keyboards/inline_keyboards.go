package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Создание инлайн клавиатуры
var MainInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(ButtonNews),
	tgbotapi.NewInlineKeyboardRow(ButtonVoiceMessages),
	tgbotapi.NewInlineKeyboardRow(ButtonSendJoke),
	tgbotapi.NewInlineKeyboardRow(ButtonUserForm),
)

// Создание клавиатуры для голосового режима
var InlineKeyboardVoiceMode = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(ButtonExitVoiceMode),
)

// Создание инлайн клавиатуры для новостей
var InlineKeyboardNews = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(ButtonTestNews),
	tgbotapi.NewInlineKeyboardRow(ButtonGameNews),
	tgbotapi.NewInlineKeyboardRow(ButtonIINews),
	tgbotapi.NewInlineKeyboardRow(ButtonHome),
)

// Инлайн кнопки для новостей
var (
	ButtonNews     = tgbotapi.NewInlineKeyboardButtonData("Новости из мира IT 📺", "Новости из мира IT")
	ButtonTestNews = tgbotapi.NewInlineKeyboardButtonData("Тестирование 🤬", "Тестирование")
	ButtonGameNews = tgbotapi.NewInlineKeyboardButtonData("Из мира игр 🎮", "Из мира игр")
	ButtonIINews   = tgbotapi.NewInlineKeyboardButtonData("Искусcтвенный интеллект 🦾", "Искусственный интеллект")
)

// Остальные инлайн кнопки
var (
	ButtonVoiceMessages = tgbotapi.NewInlineKeyboardButtonData("Голосовые сообщения 🎙", "Голосовые сообщения")
	ButtonExitVoiceMode = tgbotapi.NewInlineKeyboardButtonData("Выйти из голоcового режима ⬅", "Выйти из голосового режима")
	ButtonUserForm      = tgbotapi.NewInlineKeyboardButtonData("Пройти интервью 🤠", "Пройти интервью")
	ButtonHome          = tgbotapi.NewInlineKeyboardButtonData("На главную 🏠", "На главную")
	ButtonSendJoke      = tgbotapi.NewInlineKeyboardButtonData("Хочу шутку 🙊", "Хочу шутку")
)
