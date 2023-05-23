package main

func isPrime(value int) bool {
	prime := true
	if value == 1 || value == 0 {
		return false
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
