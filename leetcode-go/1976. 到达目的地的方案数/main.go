package main

func countPaths(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e18 // 为没有边的路径设置为 正无穷
		}
	}
	// 边带权
	for _, r := range roads {
		v, w, wt := r[0], r[1], r[2]
		g[v][w] = wt
		g[w][v] = wt
	}

	// d[i] 表示从 0 到 i 的最短路
	d := make([]int, n)
	for i := range d {
		d[i] = 1e18
	}
	d[0] = 0
	used := make([]bool, n)
	for {
		v := -1 // 维护最短的边
		for w, u := range used {
			if !u && (v < 0 || d[w] < d[v]) {
				v = w
			}
		}
		if v < 0 {
			break
		}
		used[v] = true
		for w, wt := range g[v] {
			if newD := d[v] + wt; newD < d[w] {
				d[w] = newD
			}
		}
	}

	// 只有在 DAG(有向无环图) 上才可以拓扑排序
	// 最短路构成了一个 DAG, 不需要建新图, 直接根据距离来判断每条边是否在 DAG 上
	// 计算 DAG 的入度数组
	deg := make([]int, n)
	for v, r := range g {
		for w, wt := range r {
			if d[v]+wt == d[w] { // 这条边在 DAG 上
				deg[w]++
			}
		}
	}

	// 在 DAG 上跑拓扑排序
	dp := make([]int, n) // dp[i] 表示 0 到 i 的最短路个数
	dp[0] = 1
	q := []int{0}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for w, wt := range g[v] {
			if d[v]+wt == d[w] { // 边在 DAG 上
				dp[w] = (dp[w] + dp[v]) % (1e9 + 7)
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
	}
	return dp[n-1]
}

func countPaths1(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e18 // 初始化为 +无穷
		}
	}
	for _, r := range roads {
		a, b, w := r[0], r[1], r[2]
		g[a][b] = w
		g[b][a] = w
	}
	// dijkstra 求 0 到 所有点的最短路径
	d := make([]int, n)
	for i := range d {
		d[i] = 1e18
	}
	d[0] = 0
	st := make([]bool, n)
	// yxc 模版
	// n-1轮(0 不需要), 每轮标记一个点
	// for i := 0; i < n-1; i++ {
	// 	t := -1
	// 	for j := 0; j <= n-1; j++ {
	// 		if !st[j] && (t == -1 || d[t] > d[j]) {
	// 			t = j
	// 		}
	// 	}
	// 	// 用全局最小值点 t 更新其他节点
	// 	for j := 0; j <= n-1; j++ {
	// 		d[j] = min(d[j], d[t]+g[t][j])
	// 	}
	// 	st[t] = true
	// }

	// 0x3f 模版 (适用于n 个点, 从 0 开始到 n-1)
	for {
		t := -1
		for i, flag := range st {
			if !flag && (t < 0 || d[i] < d[t]) {
				t = i
			}
		}
		if t < 0 { // 都被标记了
			break
		}
		// 遍历 t 的所有出边, 用全局最短 d[t] 更新 d[i], 0->t->i
		// t->i = d[t]+w
		for i, w := range g[t] {
			if res := d[t] + w; res < d[i] {
				d[i] = res
			}
		}
		st[t] = true
	}
	// 最短路构成了一个 DAG, 不需要建新图, 直接根据距离来判断每条边是否在 DAG 上
	// 计算 DAG 的入度数组
	deg := make([]int, n)
	for i, v := range g {
		for j, weight := range v {
			if d[i]+weight == d[j] {
				deg[j]++
			}
		}
	}
	// 在 DAG 上跑拓扑排序
	dp := make([]int, n) // dp[i] 表示 0 到 i 的最短路个数
	dp[0] = 1
	q := []int{0}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for i, weight := range g[v] {
			if d[v]+weight == d[i] { // 边在 DAG 上
				dp[i] = (dp[i] + dp[v]) % (1e9 + 7)
				// 将 i 这个点与 v 的边断开之后, 如果入度为 0, 则入队
				if deg[i]--; deg[i] == 0 {
					q = append(q, i)
				}
			}
		}
	}
	return dp[n-1]

}

func countPaths2(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e18
		}
	}
	for _, v := range roads {
		a, b, w := v[0], v[1], v[2]
		g[a][b] = w
		g[b][a] = w
	}

	// dijkstra 求 0 到 i 的最短路径
	d := make([]int, n)
	for i := range d {
		d[i] = 1e18 // 要求最小值, 初始化为最大值
	}
	d[0] = 0
	st := make([]bool, n)
	for {
		t := -1
		for i, flag := range st {
			if !flag && (t < 0 || d[i] < d[t]) {
				t = i
			}
		}
		if t < 0 {
			break
		}
		for i, w := range g[t] {
			if res := d[t] + w; res < d[i] {
				d[i] = res
			}
		}
		st[t] = true
	}

	// 最短路构成 DAG, 先求拓扑排序
	deg := make([]int, n)
	for i, v := range g {
		for j, weight := range v {
			if d[i]+weight == d[j] {
				deg[j]++
			}
		}
	}
	// 在 DEG 上做拓扑排序 (找一个顺序, 用来 dp)
	dp := make([]int, n)
	dp[0] = 1
	q := []int{0}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		for j, weight := range g[t] {
			if d[t]+weight == d[j] {
				dp[j] = (dp[j] + dp[t]) % (1e9 + 7)
				if deg[j]--; deg[j] == 0 {
					q = append(q, j)
				}
			}
		}
	}
	return dp[n-1]
}
