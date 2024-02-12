package main

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) (ans [][]int) {
	type data struct {
		col, row, val int
	}
	nodes := []data{}
	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}
		nodes = append(nodes, data{col, row, root.Val})
		dfs(root.Left, row+1, col-1)
		dfs(root.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.val < b.val && a.row == b.row)
	})
	lastCol := math.MinInt32
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return
}
