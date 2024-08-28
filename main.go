package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
	// "io"
	// "os"
)

func checkMotion(motion string) bool {

	if motion == "rock" || motion == "paper" || motion == "scissors" {

		return true
	}

	return false

}

func reading() (string, error) {
	var motion string

	_, err := fmt.Scan(&motion)

	if err != nil {
		return "", fmt.Errorf("can't reading motion: %w", err)
	}

	return strings.ToLower(motion), nil

}

func createBotMotion() string {
	motions := []string{"paper", "scissors", "rock"}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := random.Intn(3)

	return motions[randomIndex]

}

func main() {
	end := false

	moves := map[string]string{
		"paper":    "rock",
		"rock":     "scissors",
		"scissors": "paper",
	}
	for {
		userMotion, _ := reading()
		botMotion := createBotMotion()

		for !checkMotion(userMotion) {
			fmt.Println("Please, write correct item: paper, rock or scissors")
			userMotion, _ = reading()
		}

		if userMotion == botMotion {
			fmt.Print("Oops, it is draw")
		} else if moves[userMotion] == botMotion {
			fmt.Print("You win!!!")
			end = true
		} else {
			fmt.Print("Oh no, you lose :(")
		}

		fmt.Printf(" | You - [%s], Bot - [%s]\n", userMotion, botMotion)
		if !end {
			fmt.Println("Try again")
		} else {
			break
		}
	}

}
