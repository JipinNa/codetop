package leetcode

func deleteDuplicates1(head *ListNode) *ListNode {
	p := head
	for p != nil {
		if p.Next == nil {
			break
		}
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return head
}
