package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两个链表的第一个公共节点
// 方法一：哈希表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// set
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 方法二：双指针
// 本质上就是走相同的路程，如果有交点，那么一定会相遇
// 前面多走的后面少走，后面多走的前面少走，最后一定会相遇
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b== nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

