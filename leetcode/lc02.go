package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	step := false
	for l1 != nil || l2 != nil || step {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if step {
			step = false
			sum++
		}
		if sum >= 10 {
			sum -= 10
			step = true
		}
		p.Next = &ListNode{Val: sum}
		p = p.Next
	}
	return dummy.Next
}
