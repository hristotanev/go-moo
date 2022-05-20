package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const NUM_LEN = 4
const MIN, MAX = 1000, 10000

func min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}

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

	if len(guess.number) != NUM_LEN {
		return errorMessage
	}

	_, err := strconv.Atoi(guess.number)
	if err != nil {
		return errorMessage
	}

	return nil
}

func (guess *Guess) calculateNumberOfBullsAndCows(numberToGuess string) {
	guessDigits := make([]int, 10)
	numberToGuessDigits := make([]int, 10)

	for i := 0; i < len(numberToGuess); i++ {
		if guess.number[i] == numberToGuess[i] {
			guess.bulls++
		} else {
			guessDigits[guess.number[i]-'0']++
			numberToGuessDigits[numberToGuess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		guess.cows += min(numberToGuessDigits[i], guessDigits[i])
	}
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

			if guess.bulls == NUM_LEN {
				game.displaySummary()
				break
			}
		}
	}
}
