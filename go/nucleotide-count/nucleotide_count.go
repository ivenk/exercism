package dna

import (
	"strings"
	"errors"
)

type Histogram map[rune] int

type DNA string

var nucl = [4] rune {'A', 'C', 'G', 'T'}

func (d DNA) Counts() (Histogram, error) {
	var h = make(Histogram)
	// total nucleotide count found
	var nuclCount int

	 for _, n := range nucl {
		count := strings.Count(string(d), string(n) )
		//count maybe 0
		h [n] = count
		nuclCount += count
	}

	// compare the total length of the dna with the amount of valid nucleotides found.
	// A mismatch tells you that the DNA contains invalid nucleotides.
	if len(d) > nuclCount {
		return nil, errors.New("invalid nucleotides found")
	}

	return h, nil
}
