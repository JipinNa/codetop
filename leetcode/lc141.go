package leetcode

func hasCycle(head *ListNode) bool {
	end := &ListNode{}
	for head != nil {
		if head == end {
			return true
		} else {
			next := head.Next
			head.Next = end
			head = next
		}
	}
	return false
}
