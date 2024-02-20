package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据 pre[0]确定根, 根据根在 inorder 中的位置, 确定左子树/右子树, 递归下去即可
func buildTree(inorder []int, postorder []int) *TreeNode {
    idxMap := map[int]int{}
    for i, v := range inorder {
        idxMap[v] = i
    }
    var build func(int, int) *TreeNode
    build = func(inorderLeft, inorderRight int) *TreeNode {
        // 无剩余节点
        if inorderLeft > inorderRight {
            return nil
        }

        // 后序遍历的末尾元素即为当前子树的根节点
        val := postorder[len(postorder)-1]
        postorder = postorder[:len(postorder)-1]
        root := &TreeNode{Val: val}

        // 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
        // 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
        inorderRootIndex := idxMap[val]
        root.Right = build(inorderRootIndex+1, inorderRight)
        root.Left = build(inorderLeft, inorderRootIndex-1)
        return root
    }
    return build(0, len(inorder)-1)
}