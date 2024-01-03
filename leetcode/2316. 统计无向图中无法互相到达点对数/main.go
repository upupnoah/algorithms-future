package _316__统计无向图中无法互相到达点对数

func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	// 建立无向图
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vis[x] = true // 标记已访问
		size := 1
		for _, y := range g[x] {
			if !vis[y] {
				size += dfs(y)
			}
		}
		return size
	}
	total := 0
	for i, b := range vis {
		if !b {
			size := dfs(i)
			ans += int64(size) * int64(total)
			total += size
		}
	}
	return
}
