package main

import "fmt"

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
	for i := 0; i < 1000; i++ {
		fmt.Println(i, ":", isPrime(i))
	}
}
