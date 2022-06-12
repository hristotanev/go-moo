package core

import (
	"fmt"

	. "moo/helpers"
)

type Game struct {
	Number               string
	TotalNumberOfGuesses int
}

func NewGame() *Game {
	return &Game{
		Number:               GenerateNumberOfLength(NUM_LENGTH),
		TotalNumberOfGuesses: 0,
	}
}

func (game *Game) DisplaySummary() {
	fmt.Printf("%d guess(es)\n\n", game.TotalNumberOfGuesses)
}
