package leetcode_300

func rotate189(nums []int, k int) {
	reverseArray(nums)
	reverseArray(nums[:k%len(nums)])
	reverseArray(nums[k%len(nums):])
}

func reverseArray(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
