package leetcode

func isSymmetric(root *TreeNode) bool {
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		}
		if left.Val == right.Val {
			return compare(left.Left, right.Right) &&
				compare(left.Right, right.Left)
		}
		return false
	}
	return compare(root.Left, root.Right)
}
