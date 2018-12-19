package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	var result []Triplet

	// we iterate over our range
	for a := min; a <= max; a++ {
		// do the expo func
		ah := a * a
		// now we look for the matching b^2 value
		for b := a; b <= max; b++ {
			bh := b * b
			for c := a; c <= max; c++ {
				ch := c * c
				if ah+bh == ch {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}

	return result
}

func Sum(p int) []Triplet {
	var result []Triplet

	for a := 1; a < p; a++ {
		ah := a * a
		for b := a; b < p; b++ {
			bh := b * b
			ch := ah + bh
			// we just want whole numbers
			chelper := math.Sqrt(float64(ch))
			if chelper == float64(int64(chelper)) {
				c := int(chelper)
				if a+b+c == p {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}

	return result
}
