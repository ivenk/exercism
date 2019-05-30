package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, key int) string {
	var buf strings.Builder
	for _, p := range plain {
		if unicode.IsLetter(p) {
			p = applyKey(p, key)
		}
		buf.WriteRune(p)
	}
	return buf.String()
}

func applyKey(r rune, key int) rune {
	mod := 0
	add := 0
	if unicode.IsUpper(r) {
		mod = 'Z' + 1
		add = 'A'
	} else { //has to be lower
		mod = 'z' + 1
		add = 'a'
	}

	res := (int(r) + key)
	if res >= mod {
		res = (res % mod) + add
	}
	return rune(res)
}
