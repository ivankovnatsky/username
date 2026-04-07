package main

import (
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
		{"", false},
	}
	for _, tt := range tests {
		if got := isAlpha(tt.input); got != tt.want {
			t.Errorf("isAlpha(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseWords(t *testing.T) {
	input := "ab\napple\nbanana\ncherry\nfig\nextraordinary\n"
	result := parseWords(input)

	for _, w := range result {
		if len(w) < MinWordLength || len(w) > MaxWordLength {
			t.Errorf("word %q has length %d, outside [%d, %d]", w, len(w), MinWordLength, MaxWordLength)
		}
		if !isAlpha(w) {
			t.Errorf("word %q is not purely alphabetic", w)
		}
	}

	// "ab" too short, "fig" too short, "extraordinary" too long
	if len(result) != 3 {
		t.Errorf("expected 3 words, got %d: %v", len(result), result)
	}
}

func TestParseWordsFiltersNonAlpha(t *testing.T) {
	input := "hello\nworld123\ngood-bye\ntest\n"
	result := parseWords(input)

	for _, w := range result {
		if !isAlpha(w) {
			t.Errorf("word %q should have been filtered out", w)
		}
	}
}

func TestEmbeddedWordList(t *testing.T) {
	if len(words) < 1000 {
		t.Errorf("expected at least 1000 filtered words, got %d", len(words))
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

	if dupes > 5 {
		t.Errorf("too many duplicate usernames: %d out of %d", dupes, n)
	}
}
