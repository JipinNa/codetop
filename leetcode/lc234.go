package leetcode

func isPalindrome(head *ListNode) bool {
	var check func(*ListNode) bool
	front := head
	check = func(node *ListNode) bool {
		if node != nil {
			if !check(node.Next) {
				return false
			}
			if node.Val != front.Val {
				return false
			}
			front = front.Next
		}
		return true
	}
	return check(head)
}
