package main

// 原, 辅助, 目标
// 移动的过程中, 柱子的属性发生了变化 (例如 辅助柱子变为了目标柱子)

// 在递归的过程中, 要假定下一层的任务能够完成
// 我们要将 n 个盘子移动到 c, 可以看成先让 n-1 个盘子移动到 b, 然后将最后一个盘子移动到 c, 最后将 n-1 个盘子从 b 移动到 c

// 不要深入递归细节
func hanota(A []int, B []int, C []int) []int {
	var dfs func(int, *[]int, *[]int, *[]int)
	dfs = func(n int, from, mid, to *[]int) {
		if n == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}
		dfs(n-1, from, to, mid) // 先将 n-1 个盘子从 from 通过 to 移动到 mid
		dfs(1, from, mid, to)   // 将 最后一个盘子移动到 to
		dfs(n-1, mid, from, to) // 将在 mid 的 n-1 个盘子通过 from 移动到 to
	}
	dfs(len(A), &A, &B, &C)
	return C
}

// 实现一个记录所有状态的汉诺塔 function