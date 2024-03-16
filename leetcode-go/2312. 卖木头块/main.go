package main

// f[i][j] 表示: 切割一块高 i 宽 j 的木块, 能够得到的最多钱数
func sellingWood(m, n int, prices [][]int) int64 {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for _, price := range prices {
		f[price[0]][price[1]] = price[2]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 对称的, 只需要枚举一半
			for k := 1; k <= j/2; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k <= i/2; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return int64(f[m][n])
}
