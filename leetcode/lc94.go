package leetcode

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		ans = append(ans, node.Val)
		order(node.Right)
	}
	order(root)
	return ans
}
