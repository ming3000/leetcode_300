package leetcode_300

func kthSmallest(root *TreeNode, k int) int {
	var ret int

	stack := make([]*TreeNode, 0)
	for k > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} //>>
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right

		ret = top.Val
		k--
	} //>
	return ret
}
