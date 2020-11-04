package leetcode_300

func searchRange(nums []int, target int) []int {
	first, last := findFirst(nums, target), findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		} //>>
	} //>
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		} //>>
	} //>
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
