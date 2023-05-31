package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"switcher-bot/env"
	"switcher-bot/keyboards"
	"switcher-bot/task"
)

func main() {
	Bot, err := tgbotapi.NewBotAPI(env.BotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)
	
	for Update := range updates {
		if Update.CallbackQuery != nil {
			// Обработка нажатия на инлайн кнопку
			callbackData := Update.CallbackQuery.Data

			Msg = tgbotapi.NewMessage(Update.CallbackQuery.Message.Chat.ID, "")

			switch callbackData {
			case "Тестирование":
				task.TestsNews(Update)
			case "Из мира игр":
				task.GameNews(Update)
			case "Искусственный интеллект":
        task.IINews(Update)
      case "Голосовые сообщения":
        task.FuncRespondOnlyVoiceMessage(Update)
      case "Анкета":  
				task.FuncCreatingUserProfile(Update)
			case "Запрос":
			}

			callbackConfig := tgbotapi.NewCallback(Update.CallbackQuery.ID, "")
			Bot.AnswerCallbackQuery(callbackConfig)
		} else if Update.Message != nil {
			Msg = tgbotapi.NewMessage(Update.Message.Chat.ID, "")

			if Update.Message.Text != "" {
				switch Update.Message.Text {
				case "/start":

					Msg.Text = "Привет🤓\nЯ бот Егора Машина😱\n" +
						"Я пока мало что умею👉👈"

					Msg.ReplyMarkup = keyboards.Keyboard
					Bot.Send(Msg)

					Msg.Text = "Выберите тему, которая Вам интересна😌"

					Msg.ReplyMarkup = keyboards.InlineKeyboard
					Bot.Send(Msg)
        
				default:
					Msg.Text = "Пожалуйста, выберите что-то из предложенных новостей!\n" +
						"Я пока не понимаю, что Вы пишите :("
					Bot.Send(Msg)
				}
			} else {
				Msg.Text = "Я умею работать только c кнопками и командой /start :("
				Bot.Send(Msg)
			}
		}
	}
}
