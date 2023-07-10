package main

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	n, m := len(matrix), len(matrix[0])
	st := make([][]bool, n)
	for i := 0; i < n; i++ {
		st[i] = make([]bool, m)
	}
	// 上右下左， 但是我们需要的是右下左上，因此初始化 d 为 1（右）
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	x, y, d := 0, 0, 1
	// 总次数
	for i := 0; i < n*m; i++ {
		res = append(res, matrix[x][y])
		st[x][y] = true
		a, b := x+dx[d], y+dy[d]
		// 如果越界或者已经访问过，就转向
		if a < 0 || a >= n || b < 0 || b >= m || st[a][b] {
			d = (d + 1) % 4
			a, b = x+dx[d], y+dy[d]
		}
		x, y = a, b
	}
	return res
}
