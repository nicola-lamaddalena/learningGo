package main

import (
	"fmt"
)

func conjecture(x int) int {
	var steps []int
	for {
		if x == 1 {
			break
		} else if x%2 == 0 {
			x = x / 2
			steps = append(steps, x)
		} else if x%2 == 1 {
			x = x*3 + 1
			steps = append(steps, x)
		}
	}
	return len(steps)
}

func main() {
	numbers := [10]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%v: %v steps.\n", numbers[i], conjecture(numbers[i]))
	}
}
