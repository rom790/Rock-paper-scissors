package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

// fmt.Print()

func checkSpecialCommand(motion string) bool {
	return specialCommands[motion] != 0
}

func checkMotion(motion string) bool {
	motion = strings.ToLower(motion)

	return (motion == "rock" || motion == "paper" || motion == "scissors")

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

func compareMotions(userMotion, botMotion string) bool {
	end := false
	if userMotion == botMotion {
		fmt.Print("Oops, it is draw")
	} else if moves[userMotion] == botMotion {
		fmt.Print("!!!You win!!!")
		end = true
	} else {
		fmt.Print("Oh no, you lose :(")
	}
	return end
}

func checkInput(userMotion *string) {
	for !checkMotion(*userMotion) && !checkSpecialCommand(*userMotion) {
		fmt.Println("Please, write correct item: paper, rock or scissors")
		*userMotion, _ = reading()
	}
}

func main() {
	end := false
	gameMod := 0

	userMotion := ""
	for {
		botMotion := createBotMotion()
		if gameMod == 0 {
			userMotion, _ = reading()
			checkInput(&userMotion)
		}

		if gameMod == 1 {
			fmt.Printf("bot - [%s]\n", botMotion)
			userMotion, _ = reading()
			checkInput(&userMotion)
		}

		if checkSpecialCommand(userMotion) {

			switch specialCommands[userMotion] {
			case 1:
				fmt.Println("Game stopped")
				os.Exit(0)
			case 2:
				fmt.Println("game mod changed")
				gameMod = 1
				continue
			case 3:
				fmt.Println("game mod changed")
				gameMod = 0
				continue
			}
		}

		end = compareMotions(userMotion, botMotion)

		fmt.Printf(" | You - [%s], Bot - [%s]\n", userMotion, botMotion)
		if !end {
			fmt.Println("Try again")
		} else {
			break
		}
	}

}
