package task

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"switcher-bot/keyboards"	
	"switcher-bot/env"
	
)

func GameNews(Update tgbotapi.Update) {
	Msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
	"The Lord of the Rings: Gollum ушла на золото\n" +
	"https://www.igromania.ru/news/125268/The_Lord_of_the_Rings_Gollum_ushla_na_zoloto.html"
	Bot.Send(Msg)

	Msg.Text = "Очередной трейлер Diablo 4 посвятили Волшебнице\n" +
	"https://www.igromania.ru/news/125349/Ocherednoy_treyler_Diablo_4_posvyatili_Volshebnice.html"
	Bot.Send(Msg)

	Msg.Text = "🤓Список интересующих тем🤓"
	Msg.ReplyMarkup = keyboards.InlineKeyboard
	Bot.Send(Msg)
}

func TestsNews(update tgbotapi.Update) {
	Msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
	"\nКак перезапускать упавшие тесты параллельно" +
	"https://habr.com/ru/companies/wrike/articles/728826/"
	Bot.Send(Msg)

	Msg.Text = "Направления для развития ручного тестировщика в 2023 году" +
		"https://habr.com/ru/articles/728446/"
	Bot.Send(Msg)

	Msg.Text = "🤓Список интересующих тем🤓"
	Msg.ReplyMarkup = keyboards.InlineKeyboard
	Bot.Send(Msg)
}

func IINews(update tgbotapi.Update) {
	Msg.Text = "😱Нашёл пару интересных статей по этой теме😱\n" +
	"Карта «Искусственный интеллект в здравоохранении России»" +
	"https://webiomed.ru/blog/karta-iskusstvennyi-intellekt-v-zdravookhranenii-rossii/"
	Bot.Send(Msg)

	Msg.Text = "8 проблем внедрения искусственного интеллекта в здравоохранении" +
		"https://webiomed.ru/blog/8-problem-vnedreniia-iskusstvennogo-intellekta-v-zdravookhranenii/"
	Bot.Send(Msg)

	Msg.Text = "🤓Список интересующих Вас тем🤓"
	Msg.ReplyMarkup = keyboards.InlineKeyboard
	Bot.Send(Msg)
}
