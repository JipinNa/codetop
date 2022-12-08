package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var cur = 0
	var remove func(node *ListNode) *ListNode
	remove = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}
		node.Next = remove(node.Next)
		cur++
		if n == cur {
			return node.Next
		}
		return node
	}
	return remove(head)
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	n = count - n
	pre := &ListNode{
		Next: head,
	}
	for cur := pre; cur.Next != nil; cur = cur.Next {
		if n == 0 {
			next := cur.Next
			if next == nil {
				cur.Next = nil
			} else {
				cur.Next = next.Next
			}
			break
		}
		n--
	}
	return pre.Next
}
