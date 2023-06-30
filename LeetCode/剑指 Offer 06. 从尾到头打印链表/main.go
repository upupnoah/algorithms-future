package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：计算链表长度，然后创建一个数组，从头遍历链表，倒着存从 n-1 ~ 0
func reversePrint(head *ListNode) []int {
	cur := head
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	cur = head
	ans := make([]int, cnt)
	for cur != nil {
		ans[cnt-1] = cur.Val
		cnt--
		cur = cur.Next
	}
	return ans
}

// 方法二：递归
var ans []int

func reversePrint2(head *ListNode) []int {
	ans = []int{} // 清空
	ans = dfs(head)
	return ans
}

func dfs(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	dfs(head.Next)
	ans = append(ans, head.Val)
	return ans
}

// 递归的另一种写法
func reversePrint3(head *ListNode) []int {
	var f func (h *ListNode) []int 
	f = func (h *ListNode) []int {
		if h == nil {
			return []int{}
		}
		return append(f(h.Next), h.Val)
	}
	return f(head)
}
