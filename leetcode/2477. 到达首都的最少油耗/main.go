package minimum_Fuel_Cost

func minimumFuelCost(roads [][]int, seat int) (ans int64) {
	g := make([][]int, len(roads)+1) // 邻接表
	for _, road := range roads {
		a, b := road[0], road[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa { // 递归子节点, 不能递归父节点
				size += dfs(y, x) // 统计子树大小
			}
		}
		if x > 0 { // 不是根节点
			ans += int64((size-1)/seat + 1) // ceil(size/seat)
		}
		return size
	}
	dfs(0, -1)
	return
}
