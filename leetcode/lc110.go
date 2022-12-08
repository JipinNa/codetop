package leetcode

func isBalanced(root *TreeNode) bool {
	var check func(node *TreeNode, h int) (int, bool)
	check = func(node *TreeNode, h int) (int, bool) {
		if node == nil {
			return h - 1, true
		}
		lh, lb := check(node.Left, h+1)
		if !lb {
			return lh, lb
		}
		rh, rb := check(node.Right, h+1)
		if !rb {
			return rh, rb
		}
		if lh > rh {
			if lh-rh == 1 {
				return lh, true
			} else {
				return rh, false
			}
		} else {
			if rh-lh <= 1 {
				return rh, true
			} else {
				return lh, false
			}
		}
	}
	_, ok := check(root, 1)
	return ok
}
