// Package strand provides utilities for working with DNA and RNA strings.
package strand

import "strings"

var dnaComplement = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA returns a given DNA string's RNA complement.
func ToRNA(dna string) string {
	var RNA strings.Builder
	for _, rune := range dna {
		RNA.WriteString(dnaComplement[rune])
	}
	return RNA.String()
}
