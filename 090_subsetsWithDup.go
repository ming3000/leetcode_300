package leetcode_300

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	dfsSubsetsWithDup(nums, 0, nil, &ret)
	return ret
}

func dfsSubsetsWithDup(nums []int, start int, path []int, ret *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ret = append(*ret, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		} //>>
		path = append(path, nums[i])
		dfsSubsetsWithDup(nums, i+1, path, ret)
		path = path[:len(path)-1]
	} //>
}
