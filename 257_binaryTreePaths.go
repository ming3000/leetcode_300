package leetcode_300

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	dfsBinaryTreePaths(root, "", &ret)
	return ret
}

func dfsBinaryTreePaths(root *TreeNode, path string, ret *[]string) {
	if root == nil {
		return
	}
	temp := path
	temp += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, temp)
	} else {
		temp += "->"
		dfsBinaryTreePaths(root.Left, temp, ret)
		dfsBinaryTreePaths(root.Right, temp, ret)
	}
}
