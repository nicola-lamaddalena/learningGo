// Prints the primes under a value given in the command line
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

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
