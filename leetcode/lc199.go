package leetcode

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	// 层次遍历
	var view func(node *TreeNode, level int)
	view = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ret) <= level {
			ret = append(ret, node.Val)
		}
		view(node.Right, level+1)
		view(node.Left, level+1)
	}
	view(root, 0)
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	var order func(node *TreeNode, level int)
	order = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ret) <= level {
			nodes := []int{node.Val}
			ret = append(ret, nodes)
		} else {
			ret[level] = append(ret[level], node.Val)
		}
		order(node.Left, level+1)
		order(node.Right, level+1)
	}
	order(root, 0)
	return ret
}
