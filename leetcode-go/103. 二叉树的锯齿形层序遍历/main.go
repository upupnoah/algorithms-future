package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)
	turn := true
	for len(q) > 0 {
		level := len(q)
		res := []int{}
        for i := 0; i < level; i++ {
				res = append(res, q[i].Val)
				if q[i].Left != nil {
					q = append(q, q[i].Left)
				}
				if q[i].Right != nil {
					q = append(q, q[i].Right)
				}
		}
        if turn {
			n := len(res)
            for i := 0; i < n/2; i++ {
				res[i], res[n-i-1] = res[n-i-1], res[i]
			}
        }
		q = q[level:]
		ans = append(ans, res)
		turn = !turn
	}
	return
}
