package leetcode_300

import "math"

func reverse(x int) int {
	var reversed int
	for x != 0 {
		reversed = reversed*10 + x%10
		x = x / 10
		if reversed <= math.MinInt32 || reversed >= math.MaxInt32 {
			return 0
		} //>>
	} // >
	return reversed
}
