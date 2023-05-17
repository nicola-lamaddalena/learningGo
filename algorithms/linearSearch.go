package main

func linearSearch(arr []int, value int) bool {
	for i := range arr {
		if i == value {
			return true
		}
	}
	return false
}
