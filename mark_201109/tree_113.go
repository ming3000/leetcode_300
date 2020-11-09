package mark_201109

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	dfsPathSum(root, sum, nil, &ret)
	return ret
}

func dfsPathSum(root *TreeNode, sum int, path []int, ret *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			temp = append(temp, root.Val)
			*ret = append(*ret, temp)
		} //>>
		return
	} //>

	path = append(path, root.Val)
	dfsPathSum(root.Left, sum-root.Val, path, ret)
	dfsPathSum(root.Right, sum-root.Val, path, ret)
	path = path[:len(path)-1]
}
