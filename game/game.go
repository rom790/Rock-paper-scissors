package game

import (
	"fmt"
	"math/rand"
	// "os"
	"strings"
	"time"
)

const (
	Win = iota
	Lose
	Draw
	Unknown
)

var specialCommands = map[string]int{
	"stop":            1,
	"checkBotMotions": 2,
	"hideBotMotions":  3,
}

var moves = map[string]string{
	"paper":    "rock",
	"rock":     "scissors",
	"scissors": "paper",
}

type GameResult struct {
	UserMotion string
	BotMotion  string
	Status     int
	ErrorMsg   string
}

func InitResults() *GameResult {
	return &GameResult{
		UserMotion: "",
		BotMotion:  "",
		Status:     Unknown,
		ErrorMsg:   "",
	}
}

func (g *GameResult) GetResults(userMotion string) {

	if checkMsg := checkInput(userMotion); checkMsg != "" {
		g.Status = Unknown
		g.ErrorMsg = checkMsg
		return
	}

	botMotion := createBotMotion()
	end := compareMotions(userMotion, botMotion)

	g.UserMotion = userMotion
	g.BotMotion = botMotion
	g.Status = end
}

func checkSpecialCommand(motion string) bool {
	return specialCommands[motion] != 0
}

func checkMotion(motion string) bool {
	motion = strings.ToLower(motion)

	return (motion == "rock" || motion == "paper" || motion == "scissors")

}

func checkInput(userMotion string) string {

	if !checkMotion(userMotion) {
		return "Please, write correct motion ([rock], [paper] or [scissors])"
	}
	return ""
}

func reading() (string, error) {
	var motion string

	_, err := fmt.Scan(&motion)

	if err != nil {
		return "", fmt.Errorf("can't reading motion: %w", err)
	}

	return motion, nil

}

func createBotMotion() string {
	motions := []string{"paper", "scissors", "rock"}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := random.Intn(3)

	return motions[randomIndex]

}

func compareMotions(userMotion, botMotion string) int {
	res := Lose
	if userMotion == botMotion {
		res = Draw
	} else if moves[userMotion] == botMotion {
		res = Win
	}
	return res
}

func printHello() {
	fmt.Println("Hello, this is rock paper scissors")
	fmt.Printf("there are %d special commands: ", len(specialCommands))
	for key := range specialCommands {
		fmt.Printf("%v ", key)
	}
	fmt.Println("\nYou have 3 basic motions: {paper}, {scissors} and {rock}")
	fmt.Println("Write your motion:")
}
