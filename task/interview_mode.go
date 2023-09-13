package task

import (
	"strings"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ModeInterview bool
	InterviewStep int
)

var interviewData struct {
	firstName string
	age       string
	city      string
}

func StartInterview(bot *tgbotapi.BotAPI, userID int64) {
	InterviewStep = 1
	ModeInterview = true

	message.SendStringMessage(bot, userID, "Привет, меня зовут Джон!\n"+
		"Давай скорее познакомимся,\nКак тебя зовут?")
}

func ExitInterview(bot *tgbotapi.BotAPI, userID int64) {
	InterviewStep = 1
	ModeInterview = false

	message.SendStringMessage(bot, userID, "Приятного было познакомиться с Вами, " + interviewData.firstName +
		"!\n" + "Ваш возраст: " + interviewData.age + " лет\n" + "Вы живёте в городе: " + interviewData.city)
		GoHomePage(bot, userID, "Вы вышли из режима интервью!\nВыберите действие:")
}

func InterviewState(bot *tgbotapi.BotAPI, userID int64, response string) {
	switch InterviewStep {
	case 1:
		if containsDigitsOrSign(response) {
			interviewData.firstName = response
			InterviewStep++
		
			message.SendStringMessage(bot, userID, "А сколько вам лет, " + interviewData.firstName + "?")
		} else {
			message.SendStringMessage(bot, userID, "Пожалуйста, введите только своё имя!")
		}
	case 2:
		if containsLettersOrSign(response) && containsAge(response){
			interviewData.age = response
			InterviewStep++
			
			message.SendStringMessage(bot, userID, "В каком горде вы живёте, " + interviewData.firstName + "?")
		} else {
			message.SendStringMessage(bot, userID, "Пожалуйста, введите только количество лет!")
		}
	case 3:
		if containsDigitsOrSign(response) {
			interviewData.city = response
			InterviewStep++
			
			ExitInterview(bot, userID)
		} else {
			message.SendStringMessage(bot, userID, "Пожалуйста, введите только Ваш город!")
		}
	}
}

func containsDigitsOrSign(s string) bool {
	for _, char := range s {
		if (char >= '0' && char <= '9') || strings.ContainsAny(string(char), "!@#$%^&*()_+-={}[]|\\:;\"'<>,.?~") {
			return false
		}
	}
	return true
}

func containsLettersOrSign (s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func containsAge(s string) bool {
	count := 0

	for _, char := range s {
		if char >= '0' && char <= '9' {
			count++
			if count > 2 {
				return false
			}
		}
	}

	return count == 1 || count == 2 
}