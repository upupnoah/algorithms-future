package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	idx := 0
	for i, v := range postorder{
		if preorder[1] == v {
			idx = i // 左子树的根节点位置
			break
		}
	}
	// 左子树的范围在: pre 1 ~ idx+1, post 0~idx
	root.Left = constructFromPrePost(preorder[1:idx+1+1], postorder[:idx+1])
	root.Right = constructFromPrePost(preorder[idx+1+1:], postorder[idx+1:len(postorder)-1])
	return root
}