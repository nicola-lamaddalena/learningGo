package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	strInput := os.Args[1]
	intInput, err := strconv.Atoi(strInput)
	if err != nil {
		fmt.Println("You need to pass a numeric value.")
		return
	}
	// Convert the Celsius value to a Fahrenheit one, print it
	fahrenheitOutput := (float64(intInput) * 9 / 5) + 32
	fmt.Printf("%v° Celsius degrees correspond to %v° Fahrenheit degrees.\n", intInput, fahrenheitOutput)
}
