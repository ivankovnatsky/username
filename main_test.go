package main

import (
	"os"
	"testing"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"hello", true},
		{"world", true},
		{"Hello", false},
		{"hello1", false},
		{"hello-world", false},
		{"hello's", false},
		{"", true},
	}
	for _, tt := range tests {
		if got := isAlpha(tt.input); got != tt.want {
			t.Errorf("isAlpha(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseWords(t *testing.T) {
	file, err := os.Open("testdata/words.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	result, err := parseWords(file)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("expected words, got none")
	}

	for _, w := range result {
		if len(w) < MinWordLength || len(w) > MaxWordLength {
			t.Errorf("word %q has length %d, outside [%d, %d]", w, len(w), MinWordLength, MaxWordLength)
		}
		if !isAlpha(w) {
			t.Errorf("word %q is not purely alphabetic", w)
		}
	}
}

func TestParseWordsRejectsTooFew(t *testing.T) {
	file, err := os.CreateTemp("", "words-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write only one valid word (need at least 2).
	file.WriteString("solo\n")
	file.Seek(0, 0)

	_, err = parseWords(file)
	if err == nil {
		t.Error("expected error for dictionary with fewer than 2 words")
	}
}

func TestPickRandomWord(t *testing.T) {
	word, err := pickRandomWord()
	if err != nil {
		t.Fatal(err)
	}

	if len(word) < MinWordLength || len(word) > MaxWordLength {
		t.Errorf("word %q length %d outside [%d, %d]", word, len(word), MinWordLength, MaxWordLength)
	}

	if !isAlpha(word) {
		t.Errorf("word %q is not purely alphabetic", word)
	}
}

func TestGenerateUserName(t *testing.T) {
	username, err := GenerateUserName()
	if err != nil {
		t.Fatal(err)
	}

	if len(username) < MinWordLength*2 {
		t.Errorf("username %q too short (len=%d)", username, len(username))
	}

	if len(username) > MaxWordLength*2 {
		t.Errorf("username %q too long (len=%d)", username, len(username))
	}

	if !isAlpha(username) {
		t.Errorf("username %q is not purely alphabetic", username)
	}
}

func TestGenerateUserNameUniqueness(t *testing.T) {
	seen := make(map[string]bool)
	dupes := 0
	n := 100

	for i := 0; i < n; i++ {
		username, err := GenerateUserName()
		if err != nil {
			t.Fatal(err)
		}
		if seen[username] {
			dupes++
		}
		seen[username] = true
	}

	// With a reasonable word list, duplicates in 100 runs should be very rare.
	if dupes > 5 {
		t.Errorf("too many duplicate usernames: %d out of %d", dupes, n)
	}
}
