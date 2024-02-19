package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		ans = append(ans, postorder(v)...)
	}
	ans = append(ans, root.Val)
	return
}

func postorder_iterator(root *Node) (ans []int) {
	if root == nil {
		return
	}
	stk := []*Node{root}
	for len(stk) > 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		ans = append(ans, node.Val) // 访问节点

		for i := range node.Children {
			child := node.Children[i]
			if child != nil {
				stk = append(stk, child)
			}
		}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return
}
