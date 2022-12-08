package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for cur != nil {
		next := cur.Next
		for i := 0; i < k && next != nil; k++ {
			cur.Next = cur
			cur = next

		}
	}
	return head
}
