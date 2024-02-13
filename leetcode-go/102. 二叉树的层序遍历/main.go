package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ans [][]int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := []*TreeNode{}
		res := []int{}
		for _, v := range q {
			res = append(res, v.Val)
			if v.Left != nil {
				t = append(t, v.Left)
			}
			if v.Right != nil {
				t = append(t, v.Right)
			}
		}
		ans = append(ans, res)
		q = t
	}
	return 
}