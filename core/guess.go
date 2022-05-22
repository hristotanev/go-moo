package core

import (
	"errors"
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

func (guess *Guess) Validate(num_len int) error {
	errorMessage := errors.New("bad guess")

	if len(guess.Number) != num_len {
		return errorMessage
	}

	_, err := strconv.Atoi(guess.Number)
	if err != nil {
		return errorMessage
	}

	return nil
}

func (guess *Guess) CalculateNumberOfBullsAndCows(num_len int, numberToGuess string) {
	guessDigits := make([]int, 10)
	numberToGuessDigits := make([]int, 10)

	for i := 0; i < num_len; i++ {
		if guess.Number[i] == numberToGuess[i] {
			guess.Bulls++
		} else {
			guessDigits[guess.Number[i]-'0']++
			numberToGuessDigits[numberToGuess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		guess.Cows += Min(numberToGuessDigits[i], guessDigits[i])
	}
}

func (guess *Guess) GetNumber() string {
	return guess.Number
}

func (guess *Guess) GetBulls() int {
	return guess.Bulls
}

func (guess *Guess) DisplaySummary() {
	fmt.Printf("bulls: %d; cows: %d\n", guess.Bulls, guess.Cows)
}
