package main

import "fmt"

func main() {
	arr := []int{3, 5, 6, 1, 16}
	value := 1
	fmt.Println(linearSearch(arr, value))

	arr2 := []int{2, 5, 6, 11, 18, 54, 93}
	fmt.Println(binarySearch(arr2, 100))

	arr3 := []int{5, 6, 93, 11, 54, 2, 18}
	bubble(arr3)
	fmt.Println(arr3)
}
