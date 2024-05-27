// go run main.go :)
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	playGame()
}

func playGame() {
	options := []string{"rock", "paper", "scissors"}
	playerChoice := getPlayerChoice(options)
	computerChoice := getComputerChoice(options)

	fmt.Println("You chose:", playerChoice)
	fmt.Println("Computer chose:", computerChoice)

	if playerChoice == computerChoice {
		fmt.Println("It's a tie!")
	} else if (playerChoice == "rock" && computerChoice == "scissors") ||
		(playerChoice == "paper" && computerChoice == "rock") ||
		(playerChoice == "scissors" && computerChoice == "paper") {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}

func getPlayerChoice(options []string) string {
	var choice string
	fmt.Print("Enter your choice (rock/paper/scissors): ")
	fmt.Scan(&choice)
	choice = normalizeChoice(choice)

	for !isValidChoice(choice, options) {
		fmt.Print("Invalid choice. Please enter rock, paper, or scissors: ")
		fmt.Scan(&choice)
		choice = normalizeChoice(choice)
	}

	return choice
}

func getComputerChoice(options []string) string {
	return options[rand.Intn(len(options))]
}

func normalizeChoice(choice string) string {
	return strings.ToLower(strings.TrimSpace(choice))
}

func isValidChoice(choice string, options []string) bool {
	for _, validChoice := range options {
		if choice == validChoice {
			return true
		}
	}
	return false
}
