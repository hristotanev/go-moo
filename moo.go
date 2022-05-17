package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const MIN, MAX = 1000, 10000

type Game struct {
	number       string
	totalGuesses int
}

func (game *Game) displaySummary() {
	fmt.Printf("%d guess(es)\n\n", game.totalGuesses)
}

func newGame() Game {
	rand.Seed(time.Now().UnixNano())

	return Game{
		number:       fmt.Sprintf("%d", rand.Intn(MAX-MIN)+MIN),
		totalGuesses: 0,
	}
}

type Guess struct {
	number string
	bulls  int
	cows   int
}

func (guess *Guess) displaySummary() {
	fmt.Printf("bulls: %d; cows: %d\n", guess.bulls, guess.cows)
}

func (guess *Guess) formatNumber() {
	guess.number = strings.TrimSpace(guess.number)
}

func (guess *Guess) validate() error {
	errorMessage := errors.New("bad guess")

	if len(guess.number) != 4 {
		return errorMessage
	}

	_, err := strconv.Atoi(guess.number)
	if err != nil {
		return errorMessage
	}

	return nil
}

func (guess *Guess) calculateNumberOfBullsAndCows(numberToGuess string) {
}

func main() {
	fmt.Println("MOO")

	for {
		fmt.Println("new game")
		game := newGame()

		for {
			fmt.Print("? ")

			var guess Guess
			fmt.Scan(&guess.number)
			guess.formatNumber()

			game.totalGuesses++

			err := guess.validate()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			guess.calculateNumberOfBullsAndCows(game.number)
			guess.displaySummary()

			if guess.bulls == 4 {
				game.displaySummary()
				break
			}
		}
	}
}
