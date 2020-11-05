package leetcode_300

func isPalindrome9(x int) bool {
	if x < 0 {
		return false
	}

	raw, reversed := x, 0
	for raw != 0 {
		reversed = reversed*10 + raw%10
		raw = raw / 10
	} //>
	return x == reversed
}
