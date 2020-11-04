package leetcode_300

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	usedNums := make(map[int]bool)
	dfsPermute(nums, usedNums, nil, &ret)
	return ret
}

func dfsPermute(nums []int, usedNums map[int]bool, path []int, ret *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	for _, v := range nums {
		if usedNums[v] {
			continue
		} //>>
		usedNums[v] = true
		path = append(path, v)
		dfsPermute(nums, usedNums, path, ret)
		usedNums[v] = false
		path = path[:len(path)-1]
	} //>
}
