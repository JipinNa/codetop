package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}

	for floor := 1; len(queue) > 0; floor++ {
		list := make([]int, 0)
		newQueue := make([]*TreeNode, 0)
		for i := range queue {
			list = append(list, queue[i].Val)
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		if floor&1 == 0 {
			for i, n := 0, len(list); i < n/2; i++ {
				list[i], list[n-1-i] = list[n-1-i], list[i]
			}
		}
		ret = append(ret, list)
		queue = newQueue
	}
	return ret
}
