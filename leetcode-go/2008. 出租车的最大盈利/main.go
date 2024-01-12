package max_taxi_earnings

// f[i] 表示到达 i 点的最大收益
// f[i] = f[i-1]
// f[i] = max(f[i], f[j]+i-j+tips) for j in [0, i-1]
func maxTaxiEarnings(n int, rides [][]int) int64 {
	f := make([]int, n+1)
	groups := make([][][2]int, n+1)
	for _, r := range rides {
		start, end, tips := r[0], r[1], r[2]
		// 按终点位置分组
		groups[end] = append(groups[end], [2]int{start, tips})
	}
	for end := 1; end <= n; end++ {
		// 不接单
		f[end] = f[end-1]
		// 接第 i 个位置利益最大的单
		for _, r := range groups[end] {
			start, tips := r[0], r[1]
			f[end] = max(f[end], f[start]+end-start+tips)
		}
	}
	return int64(f[n])
}
