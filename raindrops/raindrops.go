// Package raindrops provides an alternative to FizzBuzz.
package raindrops

import "strconv"

/* Since ordering is important here it's not possible to use a standard map
 * as iterating over it would be in random order. An array of mappings
 * maintains the expected iteration sequence.
 */
type raindropFactor struct {
	factor      int
	replacement string
}

var conversions = []raindropFactor{
	{factor: 3, replacement: "Pling"},
	{factor: 5, replacement: "Plang"},
	{factor: 7, replacement: "Plong"},
}

// Convert changes a number to a string, the contents of which depend on the
// number's factors.
func Convert(number int) string {
	var result string

	for _, conversion := range conversions {
		if number%conversion.factor == 0 {
			result += conversion.replacement
		}
	}

	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
