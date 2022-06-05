package game

import (
	"fmt"
	"regexp"
	"strings"
)

var Win int
var Loose int

func Play(word string) {

	//create dash
	blank := strings.Repeat("-", len(word))

	guessed := false
	//return true if guess contain alphabet
	isalpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	guessed_letters := []string{}
	guessed_words := []string{}
	tries := 0
	fmt.Println(Hangmanpics(tries))
	fmt.Println(blank)

	//while loop equivalent
	for !guessed && tries < 8 {
		fmt.Println("Please Enter a letter or word: ")
		var guess string
		var check bool
		fmt.Scanf("%s", &guess)
		guess = strings.ToUpper(guess)
		if len(guess) == 1 && isalpha(guess) {
			for i := 0; i < len(guessed_letters); i++ {
				if guess == guessed_letters[i] {
					check = true
				}
			}
			if check {
				fmt.Println(Hangmanpics(tries))
				fmt.Println(blank)
				fmt.Println("You already guessed that letter!")
				fmt.Print("Guessed letters: ")
				for i := 0; i < len(guessed_letters); i++ {
					fmt.Print(guessed_letters[i])

				}
				fmt.Println()

			} else if !strings.ContainsAny(word, guess) {
				guessed_letters = append(guessed_letters, guess)
				fmt.Println(guess, " is not in the word")
				tries += 1
				fmt.Println(Hangmanpics(tries))
				fmt.Println(blank)

			} else {
				guessed_letters = append(guessed_letters, guess)
				fmt.Println(guess, " is in the word")
				dash_slice := strings.Split(blank, "")
				for i := 0; i < len(word); i++ {
					if string(word[i]) == guess {
						dash_slice[i] = guess
					}
				}
				blank = strings.Join(dash_slice, "")

				count := 0
				fmt.Println(Hangmanpics(tries))
				fmt.Println(blank)
				for _, val := range dash_slice {
					if val == "-" {
						count++

					}
				}
				if count == 0 {
					guessed = true
				}

			}
		} else if len(guess) == len(word) && isalpha(guess) {
			for i := 0; i < len(guessed_words); i++ {
				if guess == guessed_words[i] {
					fmt.Println("You already guessed this word")

				}
				guessed = true
				Win++
				blank = word
			}
			if guess != word {
				fmt.Println(guess, "is not the word")
				tries += 1

				guessed_words = append(guessed_words, guess)
			} else {
				guessed = true
				blank = word

			}

		} else {
			fmt.Println("Invalid input")
			fmt.Println(Hangmanpics(tries))
			fmt.Println(blank)
		}

	}
	if guessed {
		fmt.Println("Congratulation! You won")
		Win++
	} else {
		fmt.Println("You lost")
		fmt.Println("The word was: ", word)
		Loose++
	}
	fmt.Println("You have played ", Win+Loose, " times")
	fmt.Println("You have won ", Win, " times")
	fmt.Println("You have lost ", Loose, " times")

}
