package protein

type STOP error

type ErrInvalidBase error

// lookup table for the codons
var lookUp = map[string] string {
	"AUG" : "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG" : "Tryptophan",
	"UAA" : "STOP",
	"UAG" : "STOP",
	"UGA" : "STOP",
}

func FromCodon(codon string) (string, error) {
	var value = lookUp[codon]
	if value == "STOP" {
		return "", new(STOP)
	}
	return 	lookUp[codon], nil
}

func FromRNA(rna string) ([]string, error) {
	var currentCodon string
	var result []string
	for index, value := range []rune(rna) {
		currentCodon += string(value)

		// every third letter
		if index%2 == 0 {
			value, error := FromCodon(currentCodon)
			// error found
			if error != nil {
				// we stop the lookup
				return result, error
			} else {
				// we constantly add new values to the result
				result = append(result, value)
			}
		}
 	}
	return result, nil
}
