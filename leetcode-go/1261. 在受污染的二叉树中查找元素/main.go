package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root  *TreeNode
	exist map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	var dfs func(*TreeNode, int)
	root.Val = 0
	exist := map[int]bool{0: true}
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		root.Val = val
		exist[val] = true
		dfs(root.Left, val*2+1)
		dfs(root.Right, val*2+2)
	}
	dfs(root, 0)
	return FindElements{
		root:  root,
		exist: exist,
	}
}

func (this *FindElements) Find(target int) bool {
	return this.exist[target]
}
