package mark

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	dfsSubSet(nums, 0, nil, &ret)
	return ret
}

func dfsSubSet(nums []int, start int, path []int, ret *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ret = append(*ret, temp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfsSubSet(nums, i+1, path, ret)
		path = path[:len(path)-1]
	}
}
