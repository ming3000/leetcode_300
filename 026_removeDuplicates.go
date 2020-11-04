package leetcode_300

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	left := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i]
			left++
		} //>>
	} //>
	return left
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var left int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[left] {
			left++
			nums[left] = nums[i]
		} //>>
	} //>
	return left + 1
}
