package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路 1： 哈希表
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 思路 2：数学
// 通过环形链表 I 可以知道，fast 会追上 slow 相交在环上
// 慢走了 a（总） （入口到交点的距离）= a-b
// 快比慢多在环上走了一段路，因为慢走了 a，因此快走了 2a，所以整个环的长度为 a
// 因此交点到起点的距离 = head 到起点的距离（a-b）
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
