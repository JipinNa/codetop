package leetcode

func DetectCycle(head *ListNode) *ListNode {
	return detectCycle(head)
}

func detectCycle(head *ListNode) *ListNode {
	hashMap := make(map[*ListNode]byte)
	for head != nil {
		if _, ok := hashMap[head]; ok {
			return head
		}
		hashMap[head] = 0
		head = head.Next
	}
	return nil
}
