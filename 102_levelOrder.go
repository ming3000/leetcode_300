package leetcode_300

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLevelLen := len(queue)
		curLevelList := make([]int, 0, curLevelLen)
		for i := 0; i < curLevelLen; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevelList = append(curLevelList, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			} //>>>
			if node.Right != nil {
				queue = append(queue, node.Right)
			} //>>
		} //>>
		ret = append(ret, curLevelList)
	} //>

	return ret
}
