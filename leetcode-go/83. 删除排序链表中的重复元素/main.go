package delete_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return dummy
}
