package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Celsius to Fahrenheit converter")
	fmt.Println("Insert a numeric value: ")

	// Get the input, check for error, remove trailing chars
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured")
		return
	}
	inputStr = strings.TrimSuffix(inputStr, "\r\n")

	// Convert the input to a float value, check for error
	input, err := strconv.ParseFloat(inputStr, 32)
	if err != nil {
		fmt.Println(err)
	}

	// Convert the Celsius value to a Fahrenheit one, print it
	fahrenheitOutput := (input * 9 / 5) + 32
	fmt.Println(input, "degrees Celsius correspond to", fahrenheitOutput, "degrees Fahrenheit")
}
