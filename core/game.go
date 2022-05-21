package core

import "fmt"

type Game struct {
	Number       string
	TotalGuesses int
}

func (game *Game) GetNumber() string {
	return game.Number
}

func (game *Game) AddGuess() {
	game.TotalGuesses++
}

func (game *Game) DisplaySummary() {
	fmt.Printf("%d guess(es)\n\n", game.TotalGuesses)
}
