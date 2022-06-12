package main

import (
	"flag"
	"fmt"
	"os"

	. "moo/core"
	. "moo/helpers"
)

func init() {
	flag.IntVar(&NUM_LENGTH, "n", 4, "set number's length")
	flag.Parse()
}

func main() {
	fmt.Print("MOO\nnew game\n")

	game := NewGame()
	for {
		fmt.Print("? ")

		var guess Guess
		fmt.Scan(&guess.Number)
		guess.FormatNumber()

		game.TotalNumberOfGuesses++

		err := guess.Validate()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		guess.CalculateNumberOfBullsAndCows(game.Number)
		guess.DisplaySummary()

		if guess.Bulls == NUM_LENGTH {
			game.DisplaySummary()
			os.Exit(0)
		}
	}
}
