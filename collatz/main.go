package main

import (
	"fmt"

	"example.com/collatzSteps/util"
)

func main() {
	numbers := [6]int{1000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for i := 0; i <= len(numbers)+1; i++ {
		fmt.Println("Steps", numbers[i], ":", util.Conjecture(numbers[i]))
	}
}
