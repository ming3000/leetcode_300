package leetcode_300

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		if num&1 == 1 {
			count++
		} //>>
		num = num >> 1
	} //>
	return count
}
