package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	sums := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := []*TreeNode{}
		res := 0
		for _, v := range q {
			res += v.Val
			if v.Left != nil {
				t = append(t, v.Left)
			}
			if v.Right != nil {
				t = append(t, v.Right)
			}
		}
		q = t
		sums = append(sums, res)
	}
	if len(sums) < k {
		return -1
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	return int64(sums[k-1])
}
