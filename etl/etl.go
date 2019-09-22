//Package etl provides a function that inverts a map from ints to lists of strings to a map from strings to ints
package etl

import "strings"

// Transform takes scrabble points mapped by score and maps them by lowercase letter instead
func Transform(lettersByScore map[int][]string) map[string]int {
	scoreByLetters := make(map[string]int, 26)
	for score, letters := range lettersByScore {
		for _, letter := range letters {
			scoreByLetters[strings.ToLower(letter)] = score
		}
	}
	return scoreByLetters
}
