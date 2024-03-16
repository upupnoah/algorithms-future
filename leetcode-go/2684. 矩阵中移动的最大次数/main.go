package main

func maxMoves(gird [][]int) int {
	// n, m := len(gird), len(gird[0])
	// type pair struct{ x, y int }
	// var q []pair
	// for i := 0; i < n; i++ {
	// 	q = append(q, pair{i, 0})
	// }
	// dx, dy := []int{-1, 0, 1}, []int{1, 1, 1}
	// ans := -1
	// for len(q) > 0 {
	// 	var t []pair
	// 	for _, p := range q {
	// 		x, y := p.x, p.y
	// 		for i := 0; i < 3; i++ {
	// 			nx, ny := x+dx[i], y+dy[i]
	// 			if nx < 0 || nx >= n || ny < 0 || ny >= m || gird[nx][ny] <= gird[x][y] {
	// 				continue
	// 			}
	// 			t = append(t, pair{nx, ny})
	// 		}
	// 	}

	// 	ans++
	// 	q = t
	// }
	// return ans
	m, n := len(gird), len(gird[0])
	q := make([]int, m)
	for i := range q {
		q[i] = i
	}
	vis := make([]int, m)
	for j := 0; j < n-1; j++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			for k := max(i-1, 0); k <= min(i+1, m); k++ {
				if vis[k] != j+1 && gird[k][j+1] > gird[i][j] {
					vis[k] = j + 1 // 第 k 行目前最右访问到了 j
					q = append(q, k)
				}
			}
		}
		if q == nil { // 无法继续访问
			return j
		}
	}
	return n - 1
}
