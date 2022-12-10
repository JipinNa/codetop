package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	var backtrace func(*TreeNode, []int, int, int)
	backtrace = func(node *TreeNode, path []int, cur int, curVal int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil &&
			node.Val+curVal == targetSum {
			ret = append(ret, append([]int(nil), path...))
			return
		}
		backtrace(node.Left, path[:cur], cur+1, node.Val+curVal)
		backtrace(node.Right, path[:cur], cur+1, node.Val+curVal)
	}
	backtrace(root, []int(nil), 1, 0)
	return ret
}
