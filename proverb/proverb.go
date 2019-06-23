// Package proverb provides traditional sayings
package proverb

import "fmt"

// Proverb returns the one about horses and shoes, given a set of input strings that should describe a chain of causaulity
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))
	last := len(rhyme) - 1
	if last < 0 {
		return proverb
	}

	for i := 0; i < last; i++ {
		proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	proverb[last] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return proverb
}
