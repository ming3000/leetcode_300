package mark

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	isLeft := hasPathSum(root.Left, sum-root.Val)
	isRight := hasPathSum(root.Right, sum-root.Val)
	return isLeft || isRight
}
