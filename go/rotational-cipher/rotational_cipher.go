// got inspired by the great solution of krhubert

package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, key int) string {
	return strings.Map(applyKey(key), plain)
}

func applyKey(key int) func(rune) rune { // that syntax does not look good in go
	return func(r rune) rune {
		switch {
		case unicode.IsUpper(r):
			return rune(((int(r) + key - 'A') % 26) + 'A')
		case unicode.IsLower(r):
			return rune(((int(r) + key - 'a') % 26) + 'a')
		default:
			return r
		}
	}
}
