package main

func isPrime(value int) bool {
	prime := true
	if value == 2 {
		return prime
	}
	if value == 0 || value == 1 || value%2 == 0 {
		prime = false
	}
	if value%2 == 0 {
		return false
	}

	for i := 3; i < value; i++ {
		if value%i == 0 {
			prime = false
			return prime
		}
	}
	return prime
}
