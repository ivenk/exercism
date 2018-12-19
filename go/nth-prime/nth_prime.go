package prime

// this function is pretty expensive
func Nth(number int) (p int, ok bool) {
	if number == 0 {
		return 0, false
	}

	var primeCount = 0
	var currentNumber = 2

	// until we found the correct number of primes
	for primeCount < number {
		isPrime := true

		// check if currentNumber is a prime
		for i := 2; i <= (int(currentNumber / 2)); i++ {
			// if we can divide
			if currentNumber%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeCount++
		}
		currentNumber++
	}

	return (currentNumber - 1), true
}
