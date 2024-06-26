package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const (
	MinWordLength = 4
	MaxWordLength = 8
)

var words []string

// https://github.com/dwyl/english-words
func init() {
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) >= MinWordLength && len(word) <= MaxWordLength {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if len(words) < 2 {
		fmt.Println("Not enough words of minimum length in the list. Please check your word list.")
		os.Exit(1)
	}
}

func pickRandomWord() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		return "", err
	}
	return words[n.Int64()], nil
}

// Pick two random words to format a readable username. We choose `word1word2`
// format because most sites adhere to this format, without having hyphens, any
// other special symbols and even upper letters.
func GenerateUserName() (string, error) {
	word1, err := pickRandomWord()
	if err != nil {
		return "", err
	}
	word2, err := pickRandomWord()
	if err != nil {
		return "", err
	}
	return word1 + word2, nil
}

func main() {
	username, err := GenerateUserName()
	if err != nil {
		fmt.Println("Error generating username:", err)
		return
	}
	fmt.Println(username)
}
