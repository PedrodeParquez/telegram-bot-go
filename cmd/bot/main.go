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
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				message.SendStringMessage(bot, update.Message.Chat.ID, "Привет! Я бот Егора Машина. Что бы вы хотели сделать?")
				message.SendInlineKeyboard(bot, update.Message.Chat.ID, "Выберите действие:", &keyboards.MainInlineKeyboard)
			default:
				message.SendStringMessage(bot, update.Message.Chat.ID, "Я Вас не понимаю, выберете нужную кнопку!")
			}
		}

		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data

			switch callbackData {
			case "Новости из мира IT":
				task.News(bot, update.CallbackQuery.Message.Chat.ID)
			case "Из мира игр":
				task.GameNews(bot, update.CallbackQuery.Message.Chat.ID)
			case "Тестирование":
				task.TestNews(bot, update.CallbackQuery.Message.Chat.ID)
			case "Искусственный интеллект":
				task.IINews(bot, update.CallbackQuery.Message.Chat.ID)
			case "Голосовые сообщения":
				
			case "Пройти анкету":
				
			case "На главную":
				message.SendInlineKeyboard(bot, update.CallbackQuery.Message.Chat.ID, "Выберите действие:", &keyboards.MainInlineKeyboard)
			}
		}
	}
}
