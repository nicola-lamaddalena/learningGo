package main

import (
	"fmt"
)

func isPrime(value int) bool {
	prime := true
	if value == 1 {
		return false
	}
	for i := 2; i < value; i++ {
		if value%i == 0 {
			prime = false
			return false
		}
	}
	return prime
}
func main() {
	value := 1000
	primes := make(map[string]int)
	for i := 0; i < value; i++ {
		if isPrime(i) {
			primes["even"]++
		} else {
			primes["odd"]++
		}
	}
	fmt.Println("In the first", value, "numbers there are:")
	for k, v := range primes {
		fmt.Println(k, "-", v)
	}
}
