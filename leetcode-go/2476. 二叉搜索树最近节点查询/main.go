package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST: 中序遍历即为排序后的节点列表
func closestNodes(root *TreeNode, queries []int) (ans [][]int) {
	var dfs func(*TreeNode)
	var nodes []int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	for _, v := range queries {
		minVal, maxVal := -1, -1
		idx := sort.SearchInts(nodes, v)
		// found
		if idx < len(nodes) {
			maxVal = nodes[idx]
			if nodes[idx] == v {
				minVal, maxVal = nodes[idx], nodes[idx]
				ans = append(ans, []int{minVal, maxVal})
				continue
			}
		}
		if idx != 0 {
			minVal= nodes[idx-1]
		}
		ans = append(ans, []int{minVal, maxVal})
	}
	return
}
