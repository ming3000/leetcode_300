package mark

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ret = math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		ret = maxInt(ret, nums[i])
	} //>
	return ret
}
