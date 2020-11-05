package leetcode_300

func flatten(root *TreeNode) {
	preOrderList := preOrderTraversal(root)
	for i := 1; i < len(preOrderList); i++ {
		pre, cur := preOrderList[i-1], preOrderList[i]
		pre.Left, pre.Right = nil, cur
	}
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	preList := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		preList = append(preList, node)
		if node.Right != nil {
			stack = append(stack, node.Right)
		} //>>
		if node.Left != nil {
			stack = append(stack, node.Left)
		} //>>
	} //>
	return preList
}
