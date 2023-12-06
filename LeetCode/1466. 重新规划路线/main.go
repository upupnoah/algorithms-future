package min_reorder

func minReorder(n int, connections [][]int) int {
	g := make([][][]int, n)
	for _, edge := range connections {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], []int{y, 1}) // x -> y
		g[y] = append(g[y], []int{x, 0}) // y <- x
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		res := 0
		for _, e := range g[x] {
			if e[0] == fa {
				continue
			}
			res += e[1] + dfs(e[0], x)
		}
		return res
	}
	return dfs(0, -1)

}
