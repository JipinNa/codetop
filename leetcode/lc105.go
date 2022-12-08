package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(preIndex, inIndexStart, inIndexEnd int) *TreeNode
	build = func(preIndex, inIndexStart, inIndexEnd int) *TreeNode {
		if preIndex >= len(preorder) || inIndexStart > inIndexEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[preIndex]}
		for i := inIndexStart; i <= inIndexEnd; i++ {
			if preorder[preIndex] == inorder[i] {
				root.Left = build(preIndex+1, inIndexStart, i-1)
				root.Right = build(preIndex+i+1-inIndexStart, i+1, inIndexEnd)
				break
			}
		}
		return root
	}
	return build(0, 0, len(inorder)-1)
}
