// package hamming provides utilities for working with utf8 representations of dna strings
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming distance between two DNA strings of equal length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot calculate Hamming distance between sequences of different length")
	}
	distance := 0
	for i, base := range a {
		target, _ := utf8.DecodeRuneInString(b[i:])
		if base != target {
			distance++
		}
	}
	return distance, nil
}
