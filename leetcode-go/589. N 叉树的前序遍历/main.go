package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	// 将所有以 children 为根的 preorder append 进 ans
	for _, v := range root.Children {
		ans = append(ans, preorder(v)...)
	}
	return
}

func preorder_iterator(root *Node) (ans []int) {
	if root == nil {
		return
	}
	stk := []*Node{root}
	for len(stk) > 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		ans = append(ans, node.Val) // 访问节点

		// 逆序将子节点入栈, 以保证前序遍历的顺序
		for i := len(node.Children) - 1; i >= 0; i-- {
			child := node.Children[i]
			if child != nil {
				stk = append(stk, child)
			}
		}
	}
	return
}
