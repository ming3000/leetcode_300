package mark_201109

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.Left == nil && top.Right == nil {
			stack = stack[:len(stack)-1]
			ret = append(ret, top.Val)
			continue
		} //>>
		if top.Right != nil {
			stack = append(stack, top.Right)
			top.Right = nil
		} //>>
		if top.Left != nil {
			stack = append(stack, top.Left)
			top.Left = nil
		} //>>
	} //>

	return ret
}
