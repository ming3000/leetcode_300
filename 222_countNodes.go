package leetcode_300

func countHeight(root *TreeNode) int {
	var height int
	for root != nil {
		height++
		root = root.Left
	}
	return height
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := countHeight(root.Left)
	rightHeight := countHeight(root.Right)
	if leftHeight == rightHeight {
		return (1 << leftHeight) + countNodes(root.Right)
	} else {
		return (1 << rightHeight) + countNodes(root.Left)
	}
}
