package mark

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	if nums[low] <= nums[high] {
		return nums[low]
	}
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else {
			if nums[low] < nums[mid] {
				low = mid + 1
			} else {
				high = mid
			} //>>>
		} //>>
	} //>
	return nums[0]
}
