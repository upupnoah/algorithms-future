package main

// f[i][j] 表示: 切割一块高 i 宽 j 的木块, 能够得到的最多钱数
func sellingWood(m, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int64, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}
