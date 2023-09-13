package main

import (
	"log"
	command "switcher-bot/bot_command"
	"switcher-bot/env"
	"switcher-bot/keyboards"
	"switcher-bot/message"
	"switcher-bot/task"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(env.BotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if task.ModeInterview && update.Message != nil {
			task.InterviewState(bot, update.Message.From.ID, update.Message.Text)
			continue
		}

		if task.ModeOnlyVoiceMessage && update.Message != nil && update.Message.Voice != nil {
			task.RespondOnlyVoiceMessage(bot, update.Message.From.ID)
		}

		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data
			userID := update.CallbackQuery.From.ID

			if task.ModeOnlyVoiceMessage || task.ModeInterview {
				switch callbackData {
				case "Выйти из голосового режима":
					task.ExitOnlyVoiceMessage(bot, userID)
				default:
					continue
				}
			} else {
				switch callbackData {
				case "Новости из мира IT":
					task.News(bot, userID)
				case "Из мира игр":
					task.GameNews(bot, userID)
				case "Тестирование":
					task.TestNews(bot, userID)
				case "Искусственный интеллект":
					task.IINews(bot, userID)
				case "Голосовые сообщения":
					task.StartOnlyVoiceMessage(bot, userID)
				case "На главную":
					task.GoHomePage(bot, userID, "Выберите действие:")
				case "Хочу шутку":
					task.SendJoke(bot, userID)
				case "Пройти интервью":
					task.StartInterview(bot, userID)
				}
			}
		}

		if update.Message != nil && !task.ModeOnlyVoiceMessage {
			switch update.Message.Text {
			case "/start":
				command.Start(bot, update.Message.From.ID)
			case "/removekeyboard":
				command.RemoveKeyboard(bot, update.Message.From.ID)
			case "/addkeyboard":
				command.AddKeyboard(bot, update.Message.From.ID)
			case "/help":
				command.Help(bot, update.Message.From.ID)
			case "Кнопка 1":
				keyboards.ButtonOneClick(bot, update.Message.From.ID)
			case "Кнопка 2":
				keyboards.ButtonTwoClick(bot, update.Message.From.ID)
			default:
				message.SendStringMessage(bot, update.Message.From.ID,
					"Я Вас не понимаю, выберете нужную кнопку!")
			}
		}
	}
}
