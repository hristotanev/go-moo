package core

import (
	"testing"

	. "moo/helpers"
)

func TestGuessNumberGetsFormattedCorrectly(t *testing.T) {
	guess := Guess{
		Number: "    1234     ",
	}

	guess.FormatNumber()

	if guess.Number != "1234" {
		t.Error("Number did not get formatted correctly.")
	}
}

func TestValidationFailsWhenNumberIsNotTheRightLength(t *testing.T) {
	guess := Guess{
		Number: "12",
	}

	err := guess.Validate()
	if err == nil && err.Error() != "bad guess" {
		t.Error("Guess number is not of the right length.")
	}
}

func TestValidationFailsForAnAlphaNumericGuess(t *testing.T) {
	guess := Guess{
		Number: "12aa",
	}

	err := guess.Validate()
	if err == nil && err.Error() != "bad guess" {
		t.Errorf("Guess number is not a valid number.")
	}
}

func TestValidationPassesForAFullyNumericGuessWithCorrectLength(t *testing.T) {
	NUM_LENGTH = 4

	guess := Guess{
		Number: "1234",
	}

	err := guess.Validate()
	if err != nil {
		t.Error("Guess number could not be parsed correctly.")
	}
}

func TestNumberOfBullsAndCowsGetsCalculatedCorrectly(t *testing.T) {
	NUM_LENGTH = 4

	guess := Guess{
		Number: "1234",
		Bulls:  0,
		Cows:   0,
	}

	guess.CalculateNumberOfBullsAndCows("1374")

	expectedBulls := 2
	expectedCows := 1

	if guess.Bulls != expectedBulls || guess.Cows != expectedCows {
		t.Errorf("Expected %d bulls and %d cows, got %d bulls and %d cows.", expectedBulls, expectedCows, guess.Bulls, guess.Cows)
	}
}
