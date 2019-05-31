package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct returns the result for the largest series of digits in the given string
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return 0, errors.New("span must be smaller then string length")
	}
	if span < 0 {
		return 0, errors.New("Span must be atleast 0")
	}
	ub := span
	max := -1
	for ub <= len(digits) {
		v := 1
		for _, d := range digits[(ub - span):ub] {
			if !unicode.IsDigit(d) {
				return 0, errors.New("input string may only contain digits")
			}
			v *= int(d) - '0'
		}
		if v > max {
			max = v
		}
		ub++
	}

	return max, nil
}
