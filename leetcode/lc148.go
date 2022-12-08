package leetcode

import (
	"sort"
)

func SortList(head *ListNode) *ListNode {
	return sortList(head)
}

func sortList(head *ListNode) *ListNode {
	return sortL(head, nil)
}

func sortL(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return mergeList(sortL(head, mid), sortL(mid, tail))
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return dummy.Next
}

func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := make([]*ListNode, 0)
	for ; head != nil; head = head.Next {
		list = append(list, head)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})
	pre := ListNode{Next: list[0]}
	for i := range list {
		if i+1 < len(list) {
			list[i].Next = list[i+1]
		}
	}
	return pre.Next
}
