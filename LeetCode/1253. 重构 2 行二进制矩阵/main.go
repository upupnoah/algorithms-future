package main

// 题目：https://leetcode.cn/problems/reconstruct-a-2-row-binary-matrix/
// 我的思路：如果是 2 的话，upper 和 lower 都减 1（或者先判断colsum 中有几个 2）
// 如果是 1 的话，upper 和 lower 中大的那个减 1，如果要减 1 的时候，upper 和 lower 都为 0，那么返回空数组
// 可以在一个 for 循环中同时做
// 最后还要判断 upper 和 lower 是否都为 0，如果不是，返回空数组
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			upper--
			lower--
			ans[0][i], ans[1][i] = 1, 1
		} else if colsum[i] == 1 {
			if upper >= lower {
				upper--
				ans[0][i] = 1
			} else {
				lower--
				ans[1][i] = 1
			}
		}
		if upper < 0 || lower < 0 {
			return [][]int{}
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return ans
}
