package leetcode_300

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		} //>>
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} //>>

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				} //>>>>
				for left < right && nums[right] == nums[right-1] {
					right--
				} //>>>>
				left++
				right--
			} //>>>
		} //>>
	} //>
	return ret
}
