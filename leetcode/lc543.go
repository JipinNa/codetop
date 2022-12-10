package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := getMaxLength(root.Left) + getMaxLength(root.Right)
	max = max543(max, diameterOfBinaryTree(root.Left))
	max = max543(max, diameterOfBinaryTree(root.Right))
	return max
}

func getMaxLength(node *TreeNode) int {
	if node == nil {
		return 0
	}
	maxL := getMaxLength(node.Left)
	maxL = max543(maxL, getMaxLength(node.Right))
	return maxL + 1
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}
