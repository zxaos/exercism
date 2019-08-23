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

func ToRomanNumeral(arabic int) (string, error) {
	println(arabic)
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("can't convert ranges outside 1-3000")
	}
	var result strings.Builder

	for base := 0; arabic > 0 && base < len(numerals); {
		// find the base 10 (1, 10, 100, 1000) starting point
		count := arabic / numerals[base].value
		switch count {
		case 1:
			result.WriteRune(numerals[base].name)
		case 2:
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
		case 3:
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
		case 4:
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base-1].name)
		case 5:
			result.WriteRune(numerals[base-1].name)
		case 6:
			result.WriteRune(numerals[base-1].name)
			result.WriteRune(numerals[base].name)
		case 7:
			result.WriteRune(numerals[base-1].name)
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
		case 8:
			result.WriteRune(numerals[base-1].name)
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base].name)
		case 9:
			result.WriteRune(numerals[base].name)
			result.WriteRune(numerals[base-2].name)
		}
		arabic -= count * numerals[base].value
		base += 2
	}
	return result.String(), nil
}
