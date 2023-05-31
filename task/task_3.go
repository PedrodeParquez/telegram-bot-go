package task

import (
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "telegrambot/cmd/bot"
)

func  FuncCreatingUserProfile(update tgbotapi.Update) {
  var flag bool 
  
  Msg.Text = "Здравствуйте меня зовут Джон!\nЯ хотел бы задать " +
    "Вам несколько вопросов, если Вы не против :)\nВы готовы?"
	Bot.Send(Msg)

  if Msg.Text = "Да" {
  Msg.Text = "Представьтесь пожалуйста"
	Bot.Send(Msg)

  Msg.Text = "Сколько Вам лет,{name}?"
	Bot.Send(Msg)

  Msg.Text = "Где вы живёте, в свои {age} лет"  
  Bot.Send(Msg)

  Msg.Text = "Итак подытожим\nВас зовут {name}\nВам {}"
  }
}