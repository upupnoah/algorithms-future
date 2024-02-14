package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)
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
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return
}
