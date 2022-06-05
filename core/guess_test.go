package core

import "testing"

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

	err := guess.Validate(4)
	if err == nil && err.Error() != "bad guess" {
		t.Error("Guess number is not of the right length.")
	}
}

func TestValidationFailsForAnAlphaNumericGuess(t *testing.T) {
	guess := Guess{
		Number: "12aa",
	}

	err := guess.Validate(4)
	if err == nil && err.Error() != "bad guess" {
		t.Errorf("Guess number is not a valid number.")
	}
}

func TestValidationPassesForAFullyNumericGuessWithCorrectLength(t *testing.T) {
	guess := Guess{
		Number: "1234",
	}

	err := guess.Validate(4)
	if err != nil {
		t.Error("Guess number could not be parsed correctly.")
	}
}

func TestNumberOfBullsAndCowsGetsCalculatedCorrectly(t *testing.T) {
	guess := Guess{
		Number: "1234",
	}

	guess.CalculateNumberOfBullsAndCows(4, "1374")

	expectedBulls := 2
	expectedCows := 1

	if guess.GetBulls() != expectedBulls || guess.GetCows() != expectedCows {
		t.Errorf("Expected %d bulls and %d cows, got %d bulls and %d cows.", guess.GetBulls(), guess.GetCows(), expectedBulls, expectedCows)
	}
}
