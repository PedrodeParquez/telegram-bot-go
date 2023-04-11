package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"	
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5861728720:AAH0U6mVpIKa8kypduQysXuf8_QClQsoiF4")
	/*Данная строка используется для создания экземпляра бота Telegram API. 
	tgbotapi - это библиотека на языке Go, которая обеспечивает взаимодействие 
	с API Telegram Bot.NewBotAPI - это функция, которая принимает токен бота в 
	качестве аргумента и возвращает экземпляр BotAPI. В скобках указан API бота.*/

	if err != nil {
		log.Panic(err)
	}
	/*В данном блоке проверяется была ли ошибка при создании экземпляра 
	бота Telegram API. Если ошибка есть, то программа завершится с помощью 
	функции log.Panic, которая выведет сообщение об ошибке и прервёт выполнение
	программы. Если ошибки нет, то программа продолжит работу дальше.*/
	
	bot.Debug = true
	/*Данная строка кода устанавливает флаг отладки для экземпляра бота 
	Telegram API. Когда этот флаг установлен в значение true, бот будет 
	выводить в консоль дополнительную информацию о своей работе, такую 
	как отправленные и полученные сообщения, информацию о пользователях 
	и т.д. Это может быть полезно при отладке программы или при поиске 
	ошибок в коде. Если флаг установлен в значение false, то дополнительная 
	информация не будет выводиться.*/

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//Данная строка кода выводит в консоль имя бота, к которому мы подключились.
	
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	
	updates := bot.GetUpdatesChan(u)
	//Каждые 60 секунд в консоль будут приходить обновления о состоянии бота
	
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			//При отправлении сообщения выводит имя пользователя и текст, отправленный им.
			
			text := ""
			
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					text = "Привет, я бот Егора Машина!\nЧем я могу Вам помочь?"
				}
			} else if update.Message.Document != nil || update.Message.Photo != nil ||
				update.Message.Audio != nil || update.Message.Video != nil ||
				update.Message.Voice != nil || update.Message.Sticker != nil ||
				update.Message.VideoNote != nil || update.Message.Contact != nil ||
				update.Message.Location != nil   {
				text = "Я даже текст распознавать не могу..."
			} else if update.Message.Text != "" { 
				text = "Я не понимаю, что тут написано :("
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			//Формирование сообщения

			bot.Send(msg)
			//Отправка сообщения пользователю
		}		
	}
}

