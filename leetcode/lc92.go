package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for i := 1; i < left; i++ {
		cur = cur.Next
	}
	head = cur.Next
	for i := left; i < right; i++ {
		next := head.Next
		head.Next = next.Next
		next.Next = cur.Next
		cur.Next = next
	}

	return dummy.Next
}
