package command

import (
	"log"
	"switcher-bot/keyboards"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, userID int64) {
	message.SendUsualKeyboard(bot, userID, "Привет! Я бот Егора Машина.\n"+
		"Что бы вы хотели сделать?", &keyboards.UsualKeyboard)
	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &keyboards.MainInlineKeyboard)
}

func RemoveKeyboard(bot *tgbotapi.BotAPI, userID int64) {
	msg := tgbotapi.NewMessage(userID, "")
	msg.Text = "Клавиатура удалена! Чтобы вызвать её\n" +
		"повторно используйте команду /addkeyboard"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error removing keyboard:", err)
	}

	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &keyboards.MainInlineKeyboard)
}

func AddKeyboard(bot *tgbotapi.BotAPI, userID int64) {
	message.SendUsualKeyboard(bot, userID, "Клавиатура восстановлена!", &keyboards.UsualKeyboard)
	message.SendInlineKeyboard(bot, userID, "Выберите действие:", &keyboards.MainInlineKeyboard)
}

func Help(bot *tgbotapi.BotAPI, userID int64) {
	message.SendStringMessage(bot, userID,
		"Список всех команд данного бота 🤖\n"+
			"/start - для начала работы с ботом\n"+
			"/help - для вызова списка всех команд\n"+
			"/addkeyboard - для добавления обычной клавиатуры\n"+
			"/removekeyboard - для удаления обычной клавиатуры")
}
