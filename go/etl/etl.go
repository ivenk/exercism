package etl

import (
	"strings"
)

func Transform(input map[int][] string) map[string]int {
	var newFormat = make(map[string] int)

	// the input might contain multiple point categories
	for key, pointCategory := range input {
		// we go over each letter
		for _, letter := range pointCategory {
			newFormat[strings.ToLower(letter)] = key
		}
	}
	return newFormat
}
