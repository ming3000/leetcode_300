package leetcode_300

import "math"

func treeHeight(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(treeHeight(node.Left), treeHeight(node.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(treeHeight(root.Left)-treeHeight(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
