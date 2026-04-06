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
	var err error
	words, err = loadWords()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadWords() ([]string, error) {
	var file *os.File
	var err error

	if envPath := os.Getenv("WORD_FILE"); envPath != "" {
		file, err = os.Open(envPath)
		if err != nil {
			return nil, fmt.Errorf("error: cannot open WORD_FILE=%s: %v", envPath, err)
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
		return nil, fmt.Errorf("error: no system dictionary found. Tried: %s", strings.Join(dictPaths, ", "))
	}
	defer file.Close()

	return parseWords(file)
}

func parseWords(file *os.File) ([]string, error) {
	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if len(word) >= MinWordLength && len(word) <= MaxWordLength && isAlpha(word) {
			result = append(result, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading dictionary: %v", err)
	}

	if len(result) < 2 {
		return nil, fmt.Errorf("not enough words in the dictionary")
	}

	return result, nil
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
