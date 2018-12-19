package scrabble

import "unicode"

var onePointers = []rune {
	'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T',
}

var twoPointers = [] rune {
	'D', 'G',
}

var threePointers = [] rune {
	'B', 'C', 'M', 'P',
}

var fourPointers = [] rune {
	'F', 'H', 'V', 'W', 'Y',
}

var fivePointers = [] rune {
	'K',
}

var eightPointers = [] rune {
	'J', 'X',
}

var tenPointers = [] rune {
	'Q', 'Z',
}

func Score(word string) int {
	values := make(map[rune] int)

	// we build our map containing all the letters matched to their corresponding value
	addToValues(values, onePointers, 1)
	addToValues(values, twoPointers, 2)
	addToValues(values, threePointers, 3)
	addToValues(values, fourPointers, 4)
	addToValues(values, fivePointers, 5)
	addToValues(values, eightPointers, 8)
	addToValues(values, tenPointers, 10)

	var result int
	// now we just need to compare each letter in the given string with our value map and add the values together
	for _, x := range word {
		// there is no difference between lower and upper case letters
		// transform the letter to make it match the maps content
		result += values[unicode.ToUpper(x)]
	}

	return result
}

// function adds the given runes to the map with the value provided.
func addToValues(values map[rune] int, runes []rune ,value int) map[rune] int{
	for _, x := range runes {
		values[x] = value
	}

	return values
}
