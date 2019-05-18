package summultiples

// SumMultiples returns the sum of all multiples for the divisors given up to a certain specified limit
func SumMultiples(limit int, divisors ...int) int {
	results := make(map[int]struct{}) // this is better then bool because the empty struct does not occupy any space ( kind of like a set now )

	for _, d := range divisors { // for each divisor
		if d == 0 {
			results[0] = struct{}{}
			continue
		}
		c := d
		for c < limit { // increase till limit
			results[c] = struct{}{}
			c += d
		}
	}

	total := 0
	for r := range results {
		total += r
	}
	return total
}
