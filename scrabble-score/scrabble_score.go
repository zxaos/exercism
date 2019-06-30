// Package scrabble provides utilities for working with the game of scrabble.
package scrabble

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Score performs basic scrabble scoring given a word.
func Score(word string) int {
	var score int
	target := strings.ToLower(word)
	for _, letter := range target {
		score += scoreValues[letter]
	}
	return score
}

// ScoreSpecial performs modified scoring with per letter and total multipliers.
// It should be used when scoring words with blanks, or words placed over
// double or triple letter or word score tiles.
func ScoreSpecial(word string, letterMultipliers []int, totalMultiplier int) (int, error) {
	var score int
	target := strings.ToLower(word)
	if len(letterMultipliers) != utf8.RuneCountInString(word) {
		return 0, errors.New("Invalid per-letter multipliers")
	}

	for i, letter := range target {
		score += (scoreValues[letter] * letterMultipliers[i])
	}
	return (score * totalMultiplier), nil
}

// There's lots of ways to do this - but this is just about the simplest
// and it allow easy adjustment of score values. And it doesn't need to
// calculate a data structure at runtime.
var scoreValues = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}
