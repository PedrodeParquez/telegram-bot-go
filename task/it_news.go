package task

import (
	"switcher-bot/keyboards"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GoHomePage(bot *tgbotapi.BotAPI, userID int64, text string) {
	message.SendInlineKeyboard(bot, userID, text, &keyboards.MainInlineKeyboard)
}

func News(bot *tgbotapi.BotAPI, userID int64) {
	message.SendInlineKeyboard(bot, userID, "🤓Список интересующих статей🤓", &keyboards.InlineKeyboardNews)
}

func GameNews(bot *tgbotapi.BotAPI, userID int64) {
	msg := "😱Нашёл пару интересных статей по этой теме😱\n" +
		"The Lord of the Rings: Gollum ушла на золото\n" +
		"https://www.igromania.ru/news/125268/The_Lord_of_the_Rings_Gollum_ushla_na_zoloto.html"
	message.SendStringMessage(bot, userID, msg)

	msg = "Очередной трейлер Diablo 4 посвятили Волшебнице\n" +
		"https://www.igromania.ru/news/125349/Ocherednoy_treyler_Diablo_4_posvyatili_Volshebnice.html"
	message.SendStringMessage(bot, userID, msg)

	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &keyboards.InlineKeyboardNews)
}

func IINews(bot *tgbotapi.BotAPI, userID int64) {
	msg := "😱Нашёл пару интересных статей по этой теме😱\n" +
		"Карта «Искусственный интеллект в здравоохранении России»" +
		"https://webiomed.ru/blog/karta-iskusstvennyi-intellekt-v-zdravookhranenii-rossii/"
	message.SendStringMessage(bot, userID, msg)

	msg = "8 проблем внедрения искусственного интеллекта в здравоохранении" +
		"https://webiomed.ru/blog/8-problem-vnedreniia-iskusstvennogo-intellekta-v-zdravookhranenii/"
	message.SendStringMessage(bot, userID, msg)

	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &keyboards.InlineKeyboardNews)
}

func TestNews(bot *tgbotapi.BotAPI, userID int64) {
	msgText := "😱Нашёл пару интересных статей по этой теме😱\n" +
		"Карта «Искусственный интеллект в здравоохранении России»" +
		"https://webiomed.ru/blog/karta-iskusstvennyi-intellekt-v-zdravookhranenii-rossii/"
	message.SendStringMessage(bot, userID, msgText)

	msgText = "Направления для развития ручного тестировщика в 2023 году" +
		"https://habr.com/ru/articles/728446/"
	message.SendStringMessage(bot, userID, msgText)

	message.SendInlineKeyboard(bot, userID, "Выберите тему:", &keyboards.InlineKeyboardNews)
}



