package mark

func containsNearbyDuplicate(nums []int, k int) bool {
	checkMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if checkMap[nums[i]] {
			return true
		} //>>
		checkMap[nums[i]] = true
		if len(checkMap) > k {
			delete(checkMap, nums[i-k])
		}
	} //>
	return false
}
