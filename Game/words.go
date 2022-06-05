package game

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomWord() string {

	words := []string{}
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// append every line of txt file to word slice
	for scanner.Scan() {
		words = append(words, scanner.Text())

	}
	n := rand.Intn(len(words))
	randm := words[n]
	word := strings.ToUpper(randm)
	return word
}
