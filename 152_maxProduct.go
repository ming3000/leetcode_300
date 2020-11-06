package leetcode_300

func maxProduct(nums []int) int {
	theMax, theMin, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp := theMin
		theMin = minInt(nums[i], minInt(theMax*nums[i], theMin*nums[i]))
		theMax = maxInt(nums[i], maxInt(theMax*nums[i], temp*nums[i]))
		ret = maxInt(ret, theMax)
	}
	return ret
}
