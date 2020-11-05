package leetcode_300

func lengthOfLastWord(s string) int {
	var last int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			last++
		} else {
			if last > 0 {
				return last
			} //>>>
		} //>>
	} //>
	return last
}
