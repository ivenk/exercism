// Prints a poem
package proverb

import "fmt"

const line = "For want of a %s the %s was lost."
const ending = "And all for the want of a %s."

// Proverb prints a poem using the rhymes provided
func Proverb(rhyme []string) []string {
	var poem []string
	var lastThing string

	for _, thing := range rhyme {
		// the first item in the list is skipped
		if len(lastThing) != 0 {
			poem = append(poem, fmt.Sprintf(line, lastThing, thing))
		}
		lastThing = thing
	}

	// don't need an ending if there is no poem
	if len(rhyme) == 0 {
		return poem
	}

	return append(poem, fmt.Sprintf(ending, rhyme[0]))
}
