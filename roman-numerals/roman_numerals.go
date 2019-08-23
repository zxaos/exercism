// Package romannumerals provides utilities for creating and manipulating roman numerals
package romannumerals

import (
	"errors"
	"strings"
)

type numeral struct {
	name  rune
	value int
}

var numerals = []numeral{
	{'M', 1000},
	{'D', 500},
	{'C', 100},
	{'L', 50},
	{'X', 10},
	{'V', 5},
	{'I', 1},
}

// ToRomanNumeral returns the roman numeral representation of a natural number below 3001
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("can't convert ranges outside 1-3000")
	}

	var roman strings.Builder
	for base := 0; arabic > 0 && base < len(numerals); {
		// find the base 10 (1, 10, 100, 1000) starting point
		count := arabic / numerals[base].value

		switch count {
		case 9:
			roman.WriteRune(numerals[base].name)
			roman.WriteRune(numerals[base-2].name)
			arabic -= count * numerals[base].value
		case 4:
			roman.WriteRune(numerals[base].name)
			roman.WriteRune(numerals[base-1].name)
			arabic -= count * numerals[base].value
		case 5, 6, 7, 8: // Addition with 5
			roman.WriteRune(numerals[base-1].name)
			arabic -= numerals[base-1].value
			continue
		case 1, 2, 3: // Addition
			roman.WriteRune(numerals[base].name)
			arabic -= numerals[base].value
			continue
		}
		base += 2
	}

	return roman.String(), nil
}
