package leetcode

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		order(node.Left)
		order(node.Right)
	}
	order(root)
	return ret
}
