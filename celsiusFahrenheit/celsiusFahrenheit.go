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
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured")
		return
	}
	inputStr = strings.TrimSuffix(inputStr, "\r\n")
	input, err := strconv.ParseFloat(inputStr, 32)
	if err != nil {
		fmt.Println(err)
	}
	fahrenheitOutput := (input * 9 / 5) + 32
	fmt.Println(input, "Celsius correspond to", fahrenheitOutput, "Fahrenheit")
}
