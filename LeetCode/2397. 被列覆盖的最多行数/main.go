package _397__被列覆盖的最多行数

import "math/bits"

// 回溯: 输入视角: 选/不选
func maximumRows1(matrix [][]int, numSelect int) (ans int) {
	n, m := len(matrix), len(matrix[0]) // 行列
	selected := make([]bool, m)
	check := func() int {
		cnt := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if !selected[j] && matrix[i][j] == 1 {
					cnt--
					break
				}
			}
			cnt++
		}
		return cnt
	}
	var dfs func(int, int)
	dfs = func(i int, cnt int) {
		if i >= m { // 每个叶子有可能是答案
			if cnt == numSelect {
				ans = max(ans, check())
			}
			return
		}

		// 选第 i 列
		selected[i] = true
		dfs(i+1, cnt+1)
		selected[i] = false

		//不选第 i 列
		dfs(i+1, cnt)
	}
	dfs(0, 0)
	return
}

// 回溯: 答案视角:枚举选哪个
func maximumRows2(matrix [][]int, numSelect int) (ans int) {
	n, m := len(matrix), len(matrix[0]) // 行列
	selected := make([]bool, m)
	check := func() int {
		cnt := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if !selected[j] && matrix[i][j] == 1 {
					cnt--
					break
				}
			}
			cnt++
		}
		return cnt
	}
	var dfs func(int, int)
	dfs = func(i int, cnt int) {
		if cnt == numSelect { // 每个节点都有可能是答案
			ans = max(ans, check())
		}
		for j := i; j < m; j++ {
			selected[j] = true
			dfs(j+1, cnt+1)
			selected[j] = false
		}
	}
	dfs(0, 0)
	return
}

// 位运算
func maximumRows3(matrix [][]int, numSelect int) (ans int) {
	n, m := len(matrix), len(matrix[0]) // 行列
	mask := make([]int, n)
	for i, row := range matrix {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	// 遍历所有子集 000..000 ~ 01111..1111
	for x := 0; x < 1<<m; x++ {
		if bits.OnesCount(uint(x)) != numSelect {
			continue
		}
		cnt := 0
		for _, row := range mask {
			// 技巧: 如果 row&x == row, 说明 row 是 x 的子集(能被覆盖)
			if row&x == row { // row 是 x 的子集,所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return
}

// 位运算优化(Gosper's Hack)
func maximumRows(matrix [][]int, numSelect int) (ans int) {
	n, m := len(matrix), len(matrix[0]) // 行列
	mask := make([]int, n)
	for i, row := range matrix {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 1<<numSelect - 1; set < 1<<m; {
		cnt := 0
		for _, row := range mask {
			if row&set == row {
				cnt++
			}
		}
		ans = max(ans, cnt)
		lb := set & -set
		x := set + lb
		set = (set^x)>>bits.TrailingZeros(uint(lb))>>2 | x
	}
	return
}
