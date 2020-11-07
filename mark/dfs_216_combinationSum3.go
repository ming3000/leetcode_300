package mark

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	dfsCombinationSum3(k, n, 1, nil, &ret)
	return ret
}

func dfsCombinationSum3(size, target int, start int, path []int, ret *[][]int) {
	if target < 0 || len(path) > size {
		return
	}
	if len(path) == size && target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}

	for i := start; i <= 9; i++ {
		path = append(path, i)
		dfsCombinationSum3(size, target-i, i+1, path, ret)
		path = path[:len(path)-1]
	}
}
