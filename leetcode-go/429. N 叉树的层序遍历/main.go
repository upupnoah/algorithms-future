package main

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		t := []*Node{}
		res := []int{}
		for _, n := range  q {
			res = append(res, n.Val)
			for _, nextNode := range n.Children {
				t = append(t, nextNode)
			}
		}
		q = t
		ans = append(ans, res)
	}
	return
}