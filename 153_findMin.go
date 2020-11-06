package leetcode_300

func findMin(nums []int) int {
	var low, high, mid int
	low, high = 0, len(nums)-1

	for low <= high {
		if nums[low] <= nums[high] {
			break
		} //>>
		mid = (low + high) / 2
		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		} //>>
	} //>
	return nums[low]
}
