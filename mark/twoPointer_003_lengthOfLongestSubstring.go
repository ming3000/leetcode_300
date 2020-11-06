package mark

func lengthOfLongestSubstring(s string) int {
	checkMap := make(map[byte]int)
	var ret, slow int
	for fast := 0; fast < len(s); fast++ {
		for checkMap[s[fast]] != 0 {
			checkMap[s[slow]]--
			slow++
		} //>>
		checkMap[s[fast]] += 1
		ret = maxInt(ret, fast-slow+1)
	} //>
	return ret
}
