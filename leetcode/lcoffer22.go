package leetcode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var get func(node *ListNode)
	n := 0
	get = func(node *ListNode) {
		if node == nil {
			return
		}
		get(node.Next)
		n++
		if n == k {
			head = node
		}
	}
	get(head)
	return head
}
