package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	return rand.Intn(max - min + 1)
}

func getUserInput(guess string) string {
	fmt.Scanln(&guess)
	return guess
}

func guessGame() {
	tries := 100

	randomNumber := generateRandomNumber()

	for i := 0; i < tries; i++ {
		fmt.Println("Guess number: ")
		var guessNumber string

		guessedNumber := getUserInput(guessNumber)
		guessNumberInt, err := strconv.Atoi(guessedNumber)
		
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if guessNumberInt == randomNumber {
			fmt.Println("YOU'VE WON!!!")
			break
		} else if guessNumberInt > randomNumber {
			fmt.Println("A little bit less!")
		} else if guessNumberInt < randomNumber {
			fmt.Println("A little bit more!")
		} else {
			fmt.Println("Something went wrong")
			break
		}
	}
}

func start() {
	fmt.Println("You get 100 tries to guess the correct number!")
	guessGame()
}

func main() {
	start()
}