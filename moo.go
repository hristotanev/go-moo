package main

import (
	"flag"
	"fmt"
	"os"

	. "moo/core"
	. "moo/helpers"
)

var num_len int

func newGame() Game {
	return Game{
		Number:       GenerateNumber(num_len),
		TotalGuesses: 0,
	}
}

func init() {
	flag.IntVar(&num_len, "n", 4, "set number's length")
	flag.Parse()
}

func main() {
	fmt.Print("MOO\nnew game\n")

	game := newGame()
	for {
		fmt.Print("? ")

		var guess Guess
		fmt.Scan(&guess.Number)
		guess.FormatNumber()

		game.AddGuess()

		err := guess.Validate(num_len)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		guess.CalculateNumberOfBullsAndCows(num_len, game.GetNumber())
		guess.DisplaySummary()

		if guess.GetBulls() == num_len {
			game.DisplaySummary()
			os.Exit(0)
		}
	}
}
