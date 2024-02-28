package main

// reference: https://leetcode.cn/problems/count-number-of-possible-root-nodes/solutions/2147714/huan-gen-dppythonjavacgo-by-endlesscheng-ccwy/?envType=daily-question&envId=2024-02-29

// 若 x 与 y 之间有一条边, 那么[以 x 为根的树] 变成 [以 y 为根的数] 就只有 [x,y] 和 [y,x] 这两个猜测的正确性变了
// 因此我们可以先 dfs 一遍,算出以 0 为根的猜对次数 cnt0
// 之后再做一遍 dfs,  将 cnt0 传入, 判断 cnt 是否 >= k

// 技巧: 将这次猜对的减掉, 加上以 y 为根是否猜对(哈希表)
// reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // build graph
	}
	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))

	// guesses convert to HashMap
	for _, p := range guesses {
		s[pair{p[0], p[1]}] = 1
	}
	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时, 猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
