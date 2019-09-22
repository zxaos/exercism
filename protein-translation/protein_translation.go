// Package protein provides utilities for converting RNA base strings to proteins
package protein

import "errors"

type codon struct {
	protein string
	stop    bool
}

type protein struct {
	name   string
	codons []string
	stop   bool
}

var (
	codonLookup = make(map[string]codon)
	proteins    = []protein{
		{name: "Methionine", codons: []string{"AUG"}},
		{name: "Phenylalanine", codons: []string{"UUU", "UUC"}},
		{name: "Leucine", codons: []string{"UUA", "UUG"}},
		{name: "Serine", codons: []string{"UCU", "UCC", "UCA", "UCG"}},
		{name: "Tyrosine", codons: []string{"UAU", "UAC"}},
		{name: "Cysteine", codons: []string{"UGU", "UGC"}},
		{name: "Tryptophan", codons: []string{"UGG"}},
		{name: "", codons: []string{"UAA", "UAG", "UGA"}, stop: true}, // stop further codon processing
	}

	ErrStop        = errors.New("codon describes stop protein")
	ErrInvalidBase = errors.New("codon includes invalid base or length")
)

// FromCodon gives a protein given a valid three-base codon, or an error if invalid or a stop sequence
func FromCodon(sequence string) (protein string, err error) {
	if result, ok := codonLookup[sequence]; ok {
		if result.stop {
			err = ErrStop
		}
		protein = result.protein

	} else {
		err = ErrInvalidBase
	}
	return
}

// From RNA converts a sequence of RNA bases to a nucleotide experessed as a list of proteins
func FromRNA(rna string) (polypeptide []string, err error) {
	polypeptide = make([]string, 0, len(rna)/3)

	var protein string
	for i := 0; i < len(rna); i += 3 {
		if protein, err = FromCodon(rna[i : i+3]); err == nil {
			polypeptide = append(polypeptide, protein)
			continue
		}
		break
	}

	if err == ErrStop { // ErrStop is an "expected" error
		err = nil
	}

	return
}

func init() {
	for _, protein := range proteins {
		for _, sequence := range protein.codons {
			codonLookup[sequence] = codon{protein.name, protein.stop}
		}
	}
}
