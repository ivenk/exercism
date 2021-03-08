// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const line = "For want of a %s the %s was lost."
const ending = "And all for the want of a %s."

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	fmt.Println(len(rhyme))

	var poem []string

	for i := 0; i < len(rhyme)-1; i = i+2 {
		poem = append(poem, fmt.Sprintf(line, rhyme[i-1], rhyme[i]))
	}

	// don't need an ending if there is no poem
	if len(rhyme) == 0 {
		return poem
	}

	return append(poem, fmt.Sprintf(ending, rhyme[len(rhyme)-1]))
}
