package task

import (
	"math/rand"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FuncRespondOnlyVoiceMessage(Update tgbotapi.Update) {
	if Update.Message == nil || Update.Message.Voice == nil {
		return
	}

	variatPhrase := rand.Intn(3)
	
  text := "Hi!"
	switch variatPhrase {
	case 0:
	  text = "Я скучал по Вашему голосу!"
	case 1:
		text = "Вы что заболели?"
	case 2:
		text = "Вы сегодня прекрасно звучите!"
	}
  // message.SendStringMessage(UID, messageText)
  Bot.Send(text)
}

