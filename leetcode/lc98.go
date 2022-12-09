package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	var searchLeft func(*TreeNode, int, int) bool
	searchLeft = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min {
			return false
		}
		if node.Val >= max {
			return false
		}
		return searchLeft(node.Left, min, node.Val) &&
			searchLeft(node.Right, node.Val, max)
	}
	return searchLeft(root, math.MinInt, math.MaxInt)
}
