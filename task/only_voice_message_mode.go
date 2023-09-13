package task

import (
	"math/rand"
	"switcher-bot/keyboards"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ModeOnlyVoiceMessage bool

func StartOnlyVoiceMessage(bot *tgbotapi.BotAPI, userID int64) {
	message.SendInlineKeyboard(bot, userID, "Теперь я реагирую только на голосовые сообщения 🤯"+
		"\nЧтобы вернуться в обычный режим нажмите на кнопку",
		&keyboards.InlineKeyboardVoiceMode)
	ModeOnlyVoiceMessage = true
}

func ExitOnlyVoiceMessage(bot *tgbotapi.BotAPI, userID int64) {
	ModeOnlyVoiceMessage = false
	GoHomePage(bot, userID, "Вы вышли из голосового режима\nВыберите действие:")
}

func RespondOnlyVoiceMessage(bot *tgbotapi.BotAPI, userID int64) {
	text := ""
	variatPhrase := rand.Intn(3)

	switch variatPhrase {
	case 0:
		text = "Я скучал по Вашему голосу!"
	case 1:
		text = "Вы что заболели?"
	case 2:
		text = "Вы сегодня прекрасно звучите!"
	case 3:
		text = "Вас не слышно!"
	}

	message.SendStringMessage(bot, userID, text)
}