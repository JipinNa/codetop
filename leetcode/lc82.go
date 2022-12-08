package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	var delete func(prev, cur *ListNode, val *int) *ListNode
	delete = func(prev, cur *ListNode, val *int) *ListNode {
		if cur == nil {
			return nil
		}
		next := cur.Next
		if next != nil && cur.Val == next.Val {
			prev.Next = next
			return delete(prev, next, &next.Val)
		}
		if val != nil && cur.Val == *val {
			prev.Next = next
			if next != nil {
				return delete(prev, next, val)
			} else {
				return nil
			}
		}
		return delete(cur, next, val)
	}
	prev := &ListNode{Next: head}
	delete(prev, head, nil)
	return prev.Next
}
