package remove_nodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 逆序处理 -> 先递归, 再处理
	head.Next = removeNodes(head.Next)

	// 不是最后一个节点 && 值 < 右边节点的值 -> 返回 值大的节点
	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	}
	return head
}

func removeNodes2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeNodes2(head.Next)

	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	}
	return head
}
