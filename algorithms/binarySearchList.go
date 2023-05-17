package main

import "math"

func binarySearch(arr []int, needle int) bool {
	low := 0
	high := len(arr)
	for low < high {
		middle := low + (high-low)/2
		value := arr[int(math.Floor(float64(middle)))]
		if value == needle {
			return true
		} else if value < needle {
			low = middle + 1
		} else {
			high = middle
		}
	}
	return false
}
