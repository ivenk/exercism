// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = 0// not a triangle
    Equ Kind = 1// equilateral
    Iso Kind = 2// isosceles
    Sca Kind = 3// scalene
)

func KindFromSides(a, b, c float64) Kind {
	// test for triangle
	if a <= 0 || b <= 0 || c <= 0 || a == math.NaN() || b == math.NaN() || c == math.NaN() {
		return NaT
	} else if (a+b) < c || (b+c) < a || (a+c) < b {
		return NaT
	} else if a == b && a == c {
		//this implies b == c
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	} else {
		return Sca
	}

	// else if (a+b) == c || (a+c) == b || (b+c) == a {
	//	return
}
