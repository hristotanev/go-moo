package core

import "testing"

func TestGuessNumberGetsFormattedCorrectly(t *testing.T) {
	guess := Guess{
		Number: "    1234     ",
	}

	guess.FormatNumber()

	if guess.Number != "1234" {
		t.Fail()
	}
}

func TestValidationFailsWhenNumberIsNotTheRightLength(t *testing.T) {
	guess := Guess{
		Number: "12",
	}

	err := guess.Validate(4)
	if err == nil && err.Error() != "bad guess" {
		t.Fail()
	}
}

func TestValidationFailsForAnAlphaNumericGuess(t *testing.T) {
	guess := Guess{
		Number: "12aa",
	}

	err := guess.Validate(4)
	if err == nil && err.Error() != "bad guess" {
		t.Fail()
	}
}

func TestValidationPassesForAFullyNumericGuessWithCorrectLength(t *testing.T) {
	guess := Guess{
		Number: "1234",
	}

	err := guess.Validate(4)
	if err != nil {
		t.Fail()
	}
}
