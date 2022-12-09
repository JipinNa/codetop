package leetcode

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			sum += num*10 + node.Val
			return
		}
		dfs(node.Left, num*10+node.Val)
		dfs(node.Right, num*10+node.Val)
	}
	dfs(root, 0)
	return sum
}
