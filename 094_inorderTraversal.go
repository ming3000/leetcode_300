package leetcode_300

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} //>>
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, top.Val)
		root = top.Right
	} //>
	return ret
}
