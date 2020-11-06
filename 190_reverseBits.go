package leetcode_300

func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		var bit uint32
		if num&(1<<(32-i)) != 0 {
			bit = 1
		} else {
			bit = 0
		}
		ret += bit << i
	}
	return ret
}
