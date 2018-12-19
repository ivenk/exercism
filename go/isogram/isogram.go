package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	var inv = make(map [rune] bool)
	for _, x := range word {
		if unicode.IsLetter(x) {
			lower := unicode.ToLower(x)
			if _, cont := inv[lower]; cont {
				return false
			} else {
				// boolean variable is never read/used
				inv[lower] = false
			}
		}
	}
	return true
}
