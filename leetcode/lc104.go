package leetcode

func maxDepth(root *TreeNode) int {
	max := 0
	if root != nil {
		max = 1 + max104(maxDepth(root.Left), maxDepth(root.Right))
	}
	return max
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxDep := 0
	var dfs func(node *TreeNode, dep int)
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		maxDep = max(maxDep, dep)
		dfs(node.Left, dep+1)
		dfs(node.Right, dep+1)
	}
	dfs(root, 1)
	return maxDep
}
