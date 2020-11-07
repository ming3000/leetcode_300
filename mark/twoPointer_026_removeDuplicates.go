package mark

func removeDuplicates026(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var slow int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[slow] {
			slow++
			nums[slow] = nums[i]
		} //>>
	} //>
	return slow + 1
}
