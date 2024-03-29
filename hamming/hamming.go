// package hamming provides utilities for working with utf8 representations of dna strings
package hamming

import "errors"

// Distance calculates the Hamming distance between two DNA strings of equal length
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) != len(runesB) {
		return 0, errors.New("Cannot calculate Hamming distance between sequences of different length")
	}

	distance := 0
	for i, base := range runesA {
		if base != runesB[i] {
			distance++
		}
	}
	return distance, nil
}
