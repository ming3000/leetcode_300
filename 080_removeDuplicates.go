package leetcode_300

func removeDuplicates80First(nums []int) int {
	span := 2
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if slow < span || nums[fast] != nums[slow-span] {
			nums[slow] = nums[fast]
			slow++
		} //>>
	} //>
	return slow
}

func removeDuplicates80Second(nums []int) int {
	var slow = 1
	var count = 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count += 1
		} else {
			count = 1
		} //>>

		if count <= 2 {
			nums[slow] = nums[fast]
			slow += 1
		} //>>
	} //>

	return slow
}
