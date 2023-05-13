// Prints the primes under a value given in the command line
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

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
func main() {
	value := os.Args[1] // get the value input after the file name
	primes := make(map[string]int)

	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error - Integer value only")
		return
	}
	start := time.Now()
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
	fmt.Printf("Time elapsed %v\n", time.Since(start))
}
