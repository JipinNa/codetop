package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	count := 0
	var findPath func(*TreeNode, int)
	findPath = func(node *TreeNode, cur int) {
		if count > 0 {
			return
		}
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil &&
			cur+node.Val == targetSum {
			count++
		}
		findPath(node.Left, cur+node.Val)
		findPath(node.Right, cur+node.Val)
	}
	findPath(root, 0)
	return count > 0
}
