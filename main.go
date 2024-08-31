package main

import (
	"fmt"
	"go_l2/game"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func printHello() string {
	msg := "Hello, this is rock paper scissors"
	// fmt.Printf("there are %d special commands: ", len(specialCommands))
	// for key := range specialCommands {
	// 	fmt.Printf("%v ", key)
	// }
	msg += "\nYou have 3 basic motions: {paper}, {scissors} and {rock}"
	msg += "Write your motion:"

	return msg
}

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

		userCommand := update.Message.Command()

		var msg tgbotapi.MessageConfig

		if userCommand == "start" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, printHello())
			bot.Send(msg)
			continue
		}

		userText := strings.ToLower(update.Message.Text)

		newGame.GetResults(userText)

		if newGame.Status == game.Unknown {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, newGame.ErrorMsg)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, processResults(*newGame))
		}

		bot.Send(msg)

	}

}
