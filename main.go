package main

import (
	game "Hangman-go/Game"
	"fmt"
)

func main() {
	game.WelcomePlayer()
	word := game.GetRandomWord()
	game.Play(word)

	for {
		var choice string
		fmt.Println("Do you want to play again (y/n)?")
		fmt.Scanf("%s", &choice)
		if choice == "y" {
			word := game.GetRandomWord()
			game.Play(word)

		} else {
			break
		}

	}

}
