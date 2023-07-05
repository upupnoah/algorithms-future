package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：双指针
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	dummy := &ListNode{}
	cur := dummy
	for p1 != nil || p2 != nil {
		if p1 == nil {
			cur.Next = p2
			return dummy.Next
		}
		if p2 == nil {
			cur.Next = p1
			return dummy.Next
		}
		if p1.Val <= p2.Val {
			cur.Next = p1
			cur = cur.Next
			p1 = p1.Next
		} else {
			cur.Next = p2
			cur = cur.Next
			p2 = p2.Next
		}
	}
	return dummy.Next
}

// 方法二：递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l2.Next, l1)
		return l2
	}
}