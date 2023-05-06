package util

func Conjecture(x int) int {
	var steps []int
	for {
		if x == 1 {
			break
		} else if x%2 == 0 {
			x = x / 2
			steps = append(steps, x)
		} else if x%2 == 1 {
			x = x*3 + 1
			steps = append(steps, x)
		}
	}
	return len(steps)
}
