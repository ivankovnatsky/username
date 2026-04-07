package main

import (
	"bufio"
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	MinWordLength = 4
	MaxWordLength = 8
)

var words []string

//go:embed words_alpha.txt
var content string

func init() {
	words = parseWords(content)

	if len(words) < 2 {
		fmt.Println("Not enough words of minimum length in the list. Please check your word list.")
		os.Exit(1)
	}
}

func parseWords(input string) []string {
	var result []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if len(word) >= MinWordLength && len(word) <= MaxWordLength && isAlpha(word) {
			result = append(result, word)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading word list:", err)
		os.Exit(1)
	}
	return result
}

func isAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return len(s) > 0
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
	local := false
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--local", "-local", "-l":
			local = true
		case "--help", "-help", "-h":
			fmt.Println("Usage: username [options]")
			fmt.Println()
			fmt.Println("Generate random lowercase usernames.")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  -l, -local, --local  Generate a single-word username")
			fmt.Println("  -h, -help, --help    Show this help message")
			return
		}
	}

	var username string
	var err error
	if local {
		username, err = pickRandomWord()
	} else {
		username, err = GenerateUserName()
	}
	if err != nil {
		fmt.Println("Error generating username:", err)
		return
	}
	fmt.Println(username)
}
