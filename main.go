package main

import (
	"bufio"
	"crypto/rand"
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

// System dictionary paths, tried in order.
// Override with WORD_FILE env var (useful for NixOS).
var dictPaths = []string{
	"/usr/share/dict/words",
	"/usr/share/dict/american-english",
}

func init() {
	var file *os.File
	var err error

	if envPath := os.Getenv("WORD_FILE"); envPath != "" {
		file, err = os.Open(envPath)
		if err != nil {
			fmt.Printf("Error: cannot open WORD_FILE=%s: %v\n", envPath, err)
			os.Exit(1)
		}
	} else {
		for _, path := range dictPaths {
			file, err = os.Open(path)
			if err == nil {
				break
			}
		}
	}
	if file == nil {
		fmt.Println("Error: no system dictionary found. Tried:", strings.Join(dictPaths, ", "))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		// Only pure lowercase alpha words of the right length.
		if len(word) >= MinWordLength && len(word) <= MaxWordLength && isAlpha(word) {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading dictionary:", err)
		os.Exit(1)
	}

	if len(words) < 2 {
		fmt.Println("Not enough words in the dictionary.")
		os.Exit(1)
	}
}

func isAlpha(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
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
