package task

import (
	"math/rand"
	"switcher-bot/keyboards"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func CreatingUserProfile(bot *tgbotapi.BotAPI, userID int64, update *tgbotapi.Update) {
	msg := "Здравствуйте меня зовут Джон! 🤠\nЯ хотел бы задать " +
		"Вам несколько вопросов, если Вы не против :)\nВы готовы?"
	message.SendStringMessage(bot, userID, msg)

	switch update.Message.Text {
	case "Да":
		msg = "Представьтесь пожалуйста"
		message.SendStringMessage(bot, userID, msg)
		
		msg = "Сколько Вам лет,{name}?"
		message.SendStringMessage(bot, userID, msg)

		msg = "Где вы живёте? "
		message.SendStringMessage(bot, userID, msg)

		msg= "Итак подытожим 🤠\nВас зовут {name}\nВаш возраст: {age}\n Ваш живёте : {city}\n Рад был познакомиться с вами! 🤠"
		message.SendStringMessage(bot, userID, msg)
		message.SendInlineKeyboard(bot, userID, "Выберите действие", &keyboards.InlineKeyboardNews)
		case "Нет":
			msg = "Хорошо, если всё-таки захотите, я тут! 🤠"
			message.SendStringMessage(bot, userID, msg)
			message.SendInlineKeyboard(bot, userID, "Выберите действие", &keyboards.InlineKeyboardNews)
		default:
			msg = "Я не понимаю Вас, напишите Да или Нет 🤠"
			message.SendStringMessage(bot, userID, msg)
	}
}

var ModeOnlyVoiceMessage bool = false

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


