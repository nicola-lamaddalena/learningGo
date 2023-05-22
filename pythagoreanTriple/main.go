package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func babilonian(x int, y int) ([3]int, error) {
	if x <= y {
		return [3]int{}, errors.New("the first number must be bigger than the second one")
	}
	first := (x*x - y*y)
	second := 2 * x * y
	third := (x*x + y*y)

	triple := [3]int{first, second, third}
	return triple, nil
}

func main() {
	x := os.Args[1]
	y := os.Args[2]
	xInt, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Type integer values.")
		return
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		fmt.Println("Type integer values.")
		return
	}

	fmt.Println(babilonian(xInt, yInt))
}
