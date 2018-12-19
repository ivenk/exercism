package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	//Hamming only works for sequences of equal length
	if len(a) != len(b) {
		return 0, errors.New("sequences of unequal length do not work")
	}

	var hdistanz = 0
	// compare each bit individually
	for i:= 0; i < len(a) ; i++ {
		if a[i] != b[i] {
			hdistanz++
		}
	}

	return hdistanz, nil
}
