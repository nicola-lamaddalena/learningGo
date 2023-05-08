// Prints the primes under a given value
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(value int) bool {
	prime := true
	if value == 1 {
		return false
	}
	for i := 2; i < value; i++ {
		if value%i == 0 {
			prime = false
			return prime
		}
	}
	return prime
}
func main() {
	value := os.Args[1]
	intValue, _ := strconv.Atoi(value)
	primes := make(map[string]int)
	for i := 0; i < intValue; i++ {
		if isPrime(i) {
			primes["primes"]++
		} else {
			primes["not primes"]++
		}
	}
	fmt.Printf("In the first %v numbers there are:\n", value)
	for k, v := range primes {
		fmt.Println(k, "-", v)
	}
}
