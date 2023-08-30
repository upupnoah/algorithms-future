package _761__一个图中连通三元组的最小度数

// 思路：
// 快速计算联通三元组的度数：预处理每个点的度数,记为 degree[i]
// 对于一个联通三元组，他的度数为 degree[i] + degree[j] + degree[k] - 6

// 暴力
func minTrioDegree(n int, edges [][]int) (ans int) {
	degree := make([]int, n+1)
	g := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		g[i] = make([]int, n+1)
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x][y], g[y][x] = 1, 1
		degree[x]++
		degree[y]++
	}
	ans = 0x3f3f3f3f
	for i := 1; i < n+1; i++ {
		for j := i + 1; j < n+1; j++ {
			if g[i][j] != 1 {
				continue
			}
			for k := j + 1; k < n+1; k++ {
				if g[i][k] == 1 && g[j][k] == 1 {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == 0x3f3f3f3f {
		ans = -1
	}
	return
}

// 可以根据 degree 来构造一个有向图
// degree[i] < degree[j] （ i 指向 j
// degree[i] > degree[j] （ j 指向 i
// degree[i] = degree[j] (按照编号，小指向大
func minTrioDegree2(n int, edges [][]int) (ans int) {
	g := make([]map[int]struct{}, n+1)
	h := make([][]int, n+1)
	degree := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		g[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x][y], g[y][x] = struct{}{}, struct{}{}
		degree[x]++
		degree[y]++
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if degree[x] < degree[y] || (degree[x] == degree[y] && x < y) {
			h[x] = append(h[x], y)
		} else {
			h[y] = append(h[y], x)
		}
	}
	ans = 0x3f3f3f3f
	for i := 1; i < n+1; i++ {
		for _, j := range h[i] {
			for _, k := range h[j] {
				if _, ok := g[i][k]; ok {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == 0x3f3f3f3f {
		ans = -1
	}
	return
}
