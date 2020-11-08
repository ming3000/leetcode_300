package leetcode_300

func removeElement(nums []int, val int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		} //>>
	} //>
	return slow
}
