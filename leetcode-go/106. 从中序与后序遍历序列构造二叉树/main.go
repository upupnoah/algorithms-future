package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据 pre[0]确定根, 根据根在 inorder 中的位置, 确定左子树/右子树, 递归下去即可
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	n := len(postorder)
	rootVal := postorder[n-1]
	root := &TreeNode{rootVal, nil, nil}
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == rootVal {
            break
        }
    }
    root.Right = buildTree(inorder[i+1:], postorder[n-(len(inorder)-i-1)-1:n-1])
    root.Left = buildTree(inorder[:i], postorder[:n-(len(inorder)-i-1)-1])
    return root
}
