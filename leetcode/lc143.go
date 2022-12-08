package leetcode

func ReorderList(head *ListNode) {
	reorderList(head)
}

func reorderList(head *ListNode) {
	mid, fast := head, head
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	var right *ListNode
	for cur := mid; cur != nil; {
		next := cur.Next
		cur.Next = right
		right = cur
		cur = next
	}
	for right.Next != nil {
		rNext := right.Next
		head.Next, right.Next = right, head.Next
		right, head = rNext, right.Next
	}
}
