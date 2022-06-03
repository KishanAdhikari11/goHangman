package main

import (
	game "Hangman-go/Game"
	"fmt"
)

func main() {
	word := game.GetRandomWord()
	fmt.Println(word)

}
