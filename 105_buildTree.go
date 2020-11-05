package leetcode_300

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] {
			break
		} // If>>
	} // for>
	inorderLeft := inorder[:rootIndex]
	inorderRight := inorder[rootIndex+1:]
	preorderLeft := preorder[1 : len(inorderLeft)+1]
	preorderRight := preorder[len(inorderLeft)+1:]
	root.Left = buildTree(preorderLeft, inorderLeft)
	root.Right = buildTree(preorderRight, inorderRight)
	return root
}
