package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given array of texts concurrently
func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap, 10)
	//start concurent workers
	for _, t := range texts {
		go func(text string) {
			ch <- Frequency(text)
		}(t)
	}
	// combine results into one freqmap
	r := make(FreqMap)
	for range texts {
		for k, v := range <-ch {
			r[k] += v
		}
	}
	return r
}
