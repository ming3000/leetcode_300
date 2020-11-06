package leetcode_300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	var ret int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			} //>>>
		} //>>
		ret = maxInt(ret, dp[i])
	} //>

	return ret
}
