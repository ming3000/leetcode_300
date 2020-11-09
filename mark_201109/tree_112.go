package mark_201109

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	isLeftHas := hasPathSum(root.Left, sum-root.Val)
	isRightHas := hasPathSum(root.Right, sum-root.Val)
	return isLeftHas || isRightHas
}
