package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func SendStringMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение: %v", err)
	}
}

func SendInlineKeyboard(bot *tgbotapi.BotAPI, chatID int64, message string, inlineKeyboard *tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
	log.Printf("Не удалось отправить сообщение с инлайн клавиатурой: %v", err)
	}
}
	
func SendUsualKeyboard(bot *tgbotapi.BotAPI, chatID int64, message string, replyKeyboard *tgbotapi.ReplyKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = replyKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение с обычной клавиатурой: %v", err)
	}
}

func TextMessage( ) {

}



