package leetcode_300

func lengthOfLongestSubstring(s string) int {
	cm := make(map[byte]int)
	var left, ret int
	for i := 0; i < len(s); i++ {
		for cm[s[i]] != 0 {
			cm[s[left]]--
			left++
		} //>>
		cm[s[i]]++
		ret = maxInt(ret, i-left+1)
	} //>
	return ret
}
