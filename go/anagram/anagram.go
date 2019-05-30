package anagram

import (
	"reflect"
	"strings"
)

// Detect returns a subset of strings from the given possibilities which are actually anagrams of the given base
func Detect(base string, poss []string) []string {
	result := make([]string, 0)
	bm := buildMapFromWord(strings.ToLower(base))
	for _, p := range poss {
		lp := strings.ToLower(p)
		if lp == lbase || len(p) != len(base) { // the word itself is no anagram
			continue
		}
		pm := buildMapFromWord(lp)
		if reflect.DeepEqual(bm, pm) {
			result = append(result, p)
		}
	}
	return result
}

// buildMapFromWord returns a map matching each rune to the number of occurences in the given word.
func buildMapFromWord(word string) map[rune]int {
	m := make(map[rune]int)

	for _, l := range word {
		m[l] += 1
	}
	return m
}
