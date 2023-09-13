package task

import (
	"encoding/json"
	"log"
	"net/http"
	"switcher-bot/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Joke struct {
	Category  string `json:"category"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func GetJoke() (*Joke, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://official-joke-api.appspot.com/random_joke", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return nil, err
	}

	return &joke, nil
}

func SendJoke(bot *tgbotapi.BotAPI, userID int64) {
	joke, err := GetJoke()
	if err != nil {
		log.Println("Error getting joke:", err)
	}

	textJoke := "Шутка специально для Вас:\n\n" + joke.Setup + "\n\n" + joke.Punchline

	message.SendStringMessage(bot, userID, textJoke)
	GoHomePage(bot, userID, "Выберите действие:")
}