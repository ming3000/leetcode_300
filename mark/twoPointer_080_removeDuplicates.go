package mark

func removeDuplicates080First(nums []int) int {
	const span = 2
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if slow < span || nums[fast] != nums[slow-span] {
			nums[slow] = nums[fast]
			slow += 1
		} //>>
	} //>
	return slow
}

func removeDuplicates080Second(nums []int) int {
	var count = 1
	var slow = 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1
		} //>>
		if count <= 2 {
			nums[slow] = nums[fast]
			slow += 1
		}
	} //>

	return slow
}
