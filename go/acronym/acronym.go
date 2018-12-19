package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	var b strings.Builder
	var index int

	// we add the first letter always
	b.WriteString(strings.ToUpper(string(s[0])))   // make sure its uppercase

	// while there are still spaces or - to find we continue
	for strings.Contains(s, "-") || strings.Contains(s, " ") {
		var lineIndex = strings.Index(s, "-")
		var spaceIndex = strings.Index(s, " ")

		// check if either one has value 0
		if lineIndex < 0 {
			index = spaceIndex
		} else if spaceIndex < 0 {
			index = lineIndex
		} else {
			// if none is 0 we choose the smaller index
			if lineIndex < spaceIndex {
				index = lineIndex
			} else {
				index = spaceIndex
			}
		}

		//increase index by one to get the next letter after the seperator
		index++

		// write the next letter to the buffer and make it upper case
		b.WriteString(strings.ToUpper(string(s[index])))

		// cut off our last find
		s = s[index:]
	}


	return b.String()
}
