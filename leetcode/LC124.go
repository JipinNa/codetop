package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxSum = math.MinInt
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		maxLeft := max(maxGain(node.Left), 0)
		maxRight := max(maxGain(node.Right), 0)

		sum := node.Val + maxLeft + maxRight
		maxSum = max(sum, maxSum)
		return node.Val + max(maxLeft, maxRight)
	}
	maxGain(root)
	return maxSum
}
