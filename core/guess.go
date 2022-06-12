package core

import (
	"fmt"
	"strconv"
	"strings"

	. "moo/helpers"
)

type Guess struct {
	Number string
	Bulls  int
	Cows   int
}

func (guess *Guess) FormatNumber() {
	guess.Number = strings.TrimSpace(guess.Number)
}

func (guess *Guess) Validate() error {
	badGuessError := NewBadGuessError()

	if len(guess.Number) != NUM_LENGTH {
		return badGuessError
	}

	_, err := strconv.Atoi(guess.Number)
	if err != nil {
		return badGuessError
	}

	return nil
}

func (guess *Guess) CalculateNumberOfBullsAndCows(numberToGuess string) {
	potentialCowsOfGuessedNumber := make([]int, 10)
	potentialCowsOfNumberToGuess := make([]int, 10)

	for i := 0; i < NUM_LENGTH; i++ {
		digitOfGuessedNumber := guess.Number[i] - '0'
		digitOfNumberToGuess := numberToGuess[i] - '0'

		if digitOfGuessedNumber == digitOfNumberToGuess {
			guess.Bulls++
		} else {
			potentialCowsOfGuessedNumber[digitOfGuessedNumber]++
			potentialCowsOfNumberToGuess[digitOfNumberToGuess]++
		}
	}

	for i := 0; i < 10; i++ {
		guess.Cows += Min(potentialCowsOfNumberToGuess[i], potentialCowsOfGuessedNumber[i])
	}
}

func (guess *Guess) DisplaySummary() {
	fmt.Printf("bulls: %d; cows: %d\n", guess.Bulls, guess.Cows)
}
