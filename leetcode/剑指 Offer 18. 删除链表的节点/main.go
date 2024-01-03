package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 定义一个虚拟头结点（防止删除第一个节点越界，不想特判）
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next:head}
	pre := dummy
	// 找到要删除节点的前一个节点
	for pre.Next.Val != val {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}