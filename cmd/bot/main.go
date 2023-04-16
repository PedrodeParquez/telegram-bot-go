package main

import (
	"log"
	"telegrambot/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(env.BotToken)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg.Text = "Привет, неизвестный пользователь😃\nЯ бот Егора Машина😱\n" +
					"Выберите интересующую вас тему😌"		
				
				bot.Send(msg)

				buttons := []tgbotapi.KeyboardButton{
					tgbotapi.NewKeyboardButton("Тестирование"),
					tgbotapi.NewKeyboardButton("Из мира игр"),
					tgbotapi.NewKeyboardButton("Искуственный интеллект"),
				}

				// Добавляем кнопки в клавиатуру и прикрепляем к сообщению
				keyboard := tgbotapi.NewReplyKeyboard(buttons)
				msg.ReplyMarkup = keyboard

			case "Тестирование🤬":
				msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
					"\nКак перезапускать упавшие тесты параллельно" +
					"https://habr.com/ru/companies/wrike/articles/728826/"
				bot.Send(msg)
				
				msg.Text = "Направления для развития ручного тестировщика в 2023 году" +
					"https://habr.com/ru/articles/728446/"
				bot.Send(msg)
			case "Из мира игр 🎮":
				msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
					"The Lord of the Rings: Gollum ушла на золото\n" +
					"https://www.igromania.ru/news/125268/The_Lord_of_the_Rings_Gollum_ushla_na_zoloto.html" 
				bot.Send(msg)

				msg.Text = "Очередной трейлер Diablo 4 посвятили Волшебнице\n" +
					"https://www.igromania.ru/news/125349/Ocherednoy_treyler_Diablo_4_posvyatili_Volshebnice.html"
				bot.Send(msg)
			case "Искуственный интеллект🦾":
				msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
					"Карта «Искусственный интеллект в здравоохранении России»" +
					"https://webiomed.ru/blog/karta-iskusstvennyi-intellekt-v-zdravookhranenii-rossii/"
				bot.Send(msg)

				msg.Text = "8 проблем внедрения искусственного интеллекта в здравоохранении" +
					"https://webiomed.ru/blog/8-problem-vnedreniia-iskusstvennogo-intellekta-v-zdravookhranenii/"
				bot.Send(msg)
			default:
				msg.Text = "Пожалуйста, выберите что-то изпредложенных новостей"
				bot.Send(msg)
			}
		}
	}
}
