package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Celsius to Fahrenheit converter")
	fmt.Println("Insert a numeric value: ")

	strInput := os.Args[1]
	intInput, _ := strconv.Atoi(strInput)

	// Convert the Celsius value to a Fahrenheit one, print it
	fahrenheitOutput := (float64(intInput) * 9 / 5) + 32
	fmt.Printf("%v° Celsius degrees correspond to %v° Fahrenheit degrees.\n", intInput, fahrenheitOutput)
}
