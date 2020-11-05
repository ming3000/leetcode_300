package leetcode_300

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return nil
	}
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	ret := make([][]int, 0)
	dfsCombine(nums, k, 0, nil, &ret)
	return ret
}

func dfsCombine(nums []int, size int, start int, path []int, ret *[][]int) {
	if len(path) == size {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfsCombine(nums, size, i+1, path, ret)
		path = path[:len(path)-1]
	}
}
