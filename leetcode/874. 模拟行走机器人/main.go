package main

// 思路：模拟 + 哈希表
func robotSim(commands []int, obstacles [][]int) int {
	st := make(map[int]bool)
	for i := 0; i < len(obstacles); i++ {
		st[obstacles[i][0]*100003+obstacles[i][1]] = true
	}
	// 北东南西
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	x, y, d := 0, 0, 0
	ans := 0
	for _, c := range commands {
		if c == -1 {
			d = (d + 1) % 4
		} else if c == -2 {
			// 向左转相当于右转 3 次
			d = (d + 3) % 4
		} else {
			for i := 1; i <= c; i++ {
				a, b := x+dx[d], y+dy[d]
				if st[a*100003+b] {
					break
				}
				x, y = a, b
				ans = max(ans, x*x+y*y)

			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
