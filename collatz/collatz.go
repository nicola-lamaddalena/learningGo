package main

import "fmt"

func conjecture(x int) int {
	var collatzSteps []int
	for {
		if x == 1 {
			break
		} else if x%2 == 0 {
			x = x / 2
			collatzSteps = append(collatzSteps, x)
		} else if x%2 == 1 {
			x = x*3 + 1
			collatzSteps = append(collatzSteps, x)
		}
	}
	return len(collatzSteps)
}

func main() {
	fmt.Println(conjecture(40000))
}
