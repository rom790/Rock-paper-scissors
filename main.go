package main

import (
	"fmt"
	// "fmt"
	"go_l2/game"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// func sendMessage(text string, bot *tgbotapi.BotAPI, chatID int64){
// 	msg := tgbotapi.NewMessage(chatID, text)
// 	bot.Send(msg)
// }

// const (
// 	Win = iota
// 	Lose
// 	Draw
// 	Unknown
// )

func processResults(g game.GameResult) string {

	var msg string

	if g.Status == game.Win {
		msg = "!You win!"
	} else if g.Status == game.Lose {
		msg = "Oops, you lose"
	} else if g.Status == game.Draw {
		msg = "It is draw, try again"
	}

	msg += fmt.Sprintf(" || Your motion - [%s]  Bot motion - [%s]", g.UserMotion, g.BotMotion)

	return msg
}

func main() {

	newGame := game.InitResults()
	token := os.Getenv("TelegramBotToken")

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	channel := tgbotapi.NewUpdate(0)
	channel.Timeout = 60

	updates := bot.GetUpdatesChan(channel)

	for update := range updates {

		userText := strings.ToLower(update.Message.Text)

		var msg tgbotapi.MessageConfig

		newGame.GetResults(userText)

		if newGame.Status == game.Unknown {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, newGame.ErrorMsg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, processResults(*newGame))
		}

		bot.Send(msg)

	}

}
