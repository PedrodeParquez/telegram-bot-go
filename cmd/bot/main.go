package main

import (
	"log"
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
	
	u:= tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if task.ModeOnlyVoiceMessage && update.Message.Voice != nil {
			task.RespondOnlyVoiceMessage(bot, update.Message.From.ID)

			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case "Выйти из голосового режима":
					task.ModeOnlyVoiceMessage = false
				}
			}
		}

		if update.CallbackQuery != nil && !task.ModeOnlyVoiceMessage {
			callbackData := update.CallbackQuery.Data

			switch callbackData {
			case "Новости из мира IT":
				task.News(bot, update.CallbackQuery.Message.From.ID)
			case "Из мира игр":
				task.GameNews(bot, update.CallbackQuery.Message.From.ID)
			case "Тестирование":
				task.TestNews(bot, update.CallbackQuery.Message.From.ID)
			case "Искусственный интеллект":
				task.IINews(bot, update.CallbackQuery.Message.From.ID)
			case "Голосовые сообщения":
				message.SendInlineKeyboard(bot, update.CallbackQuery.Message.From.ID, 
				"Теперь я реагирую только на голосовые сообщения 🤯" +
				"\nЧтобы вернуться в обычный режим нажмите на кнопку",
				&keyboards.InlineKeyboardVoiceMode)
				task.ModeOnlyVoiceMessage = true		
			case "На главную":
				message.SendInlineKeyboard(bot, update.CallbackQuery.Message.From.ID, "Выберите действие:", &keyboards.MainInlineKeyboard)
			}
		}

		if update.Message != nil && !task.ModeOnlyVoiceMessage {
			switch update.Message.Text {
			case "/start":
				message.SendStringMessage(bot, update.Message.From.ID, "Привет! Я бот Егора Машина.\nЧто бы вы хотели сделать?")
				message.SendInlineKeyboard(bot, update.Message.From.ID, "Выберите действие:", &keyboards.MainInlineKeyboard)
			default:
				message.SendStringMessage(bot, update.Message.From.ID, "Я Вас не понимаю, выберете нужную кнопку!")
			}
		}
	}
}
