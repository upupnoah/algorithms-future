package main

// 前置题目 198 + 二维前缀和

// 启发思考：寻找子问题
// 设 pizza 有 n 行 m 列
// 砍一刀分 2块，砍两刀分 3 块，砍三刀分 4 块，...，砍 k-1 刀分 k 块
// 水平切：上面部分的高度为 h，还剩下 n-h 行 m 列的矩形，还需要切 k-2 刀
// 垂直切：左边部分的宽度为 w，还剩下 n 行 m-w 列的矩形，还需要切 k-2 刀
// 无论怎么切，得到的都是 和原问题相似的、规模更小的子问题，所以可以用递归解决

// 动态规划有 选/不选 和 枚举选哪个 这两种基本的思考方式
// 这题用到的是：枚举选哪个

// 递归怎么写：状态定义与状态转移方程
// 我们要处理的子矩形，右下角总是 （n-1, m-1）,因此只需要知道左上角即可确定子矩形
// 定义 dfs(c,i,j)表示把左上角在（i,j) 右下角在(m-1,n-1)的子矩形切 c 刀，每块都包含至少一个苹果的方案数
// 如果 c>=1, 分类讨论:
// 1. 水平切: 枚举 i2（切出的高度是 i2-i），剩下的子矩形左上角在 (i2,j)，还需要切 c-1 刀，即dfs(c-1,i2,j)
// 2. 垂直切: 枚举 j2（切出的宽度是 j2-j），剩下的子矩形左上角在 (i,j2)，还需要切 c-1 刀，即dfs(c-1,i,j2)
// 累加这些方案数: dfs(c,i,j) = sum(dfs(c-1,i,j2), (j<j2<n) + sum(dfs(c-1,i2,j), (i<i2<m)
// 需要判断子矩形是否有苹果

// 递归超时版本
//
//	func ways(pizza []string, k int) int {
//		const mod = 1_000_000_007
//		ms := NewMatrixSum(pizza)
//		n, m := len(pizza), len(pizza[0])
//		var dfs func(int, int, int) int
//		dfs = func(c, i, j int) (res int) {
//			if c == 0 { // 递归边界：无法再切
//				if ms.query(i, j, n, m) > 0 { // 有苹果
//					return 1 // 合法
//				}
//				return 0 // 不合法
//			}
//			// 思想：枚举选哪个
//			for i2 := i + 1; i2 < n; i2++ { // 水平切
//				if ms.query(i, j, i2, m) > 0 {
//					res += dfs(c-1, i2, j)
//				}
//			}
//			for j2 := j + 1; j2 < m; j2++ { // 垂直切
//				if ms.query(i, j, n, j2) > 0 {
//					res += dfs(c-1, i, j2)
//				}
//			}
//			return res % mod
//		}
//		return dfs(k-1, 0, 0)
//	}
//

// MatrixSum 二维前缀和模板（'A' 的 ASCII 码最低位为 1，'.' 的 ASCII 码最低位为 0）
type MatrixSum [][]int

func NewMatrixSum(matrix []string) MatrixSum {
	n, m := len(matrix), len(matrix[0])
	sum := make([][]int, n+1)
	sum[0] = make([]int, m+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, m+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + int(x&1)
		}
	}
	return sum
}

// 返回左上角在 (r1,c1) 右下角在 (r2-1,c2-1) 的子矩阵元素和（类似前缀和的左闭右开）
func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	return s[r2][c2] - s[r1][c2] - s[r2][c1] + s[r1][c1]
}

// 递归+记录返回值 = 记忆化搜索
//func ways(pizza []string, k int) int {
//	const mod = 1_000_000_007
//	ms := NewMatrixSum(pizza)
//	n, m := len(pizza), len(pizza[0])
//	memo := make([][][]int, k)
//	for c := range memo {
//		memo[c] = make([][]int, n)
//		for i := range memo[c] {
//			memo[c][i] = make([]int, m)
//			for j := range memo[c][i] {
//				memo[c][i][j] = -1 // 没有被计算过
//			}
//		}
//	}
//	var dfs func(int, int, int) int
//	dfs = func(c, i, j int) (res int) {
//		if c == 0 {
//			if ms.query(i, j, n, m) > 0 {
//				return 1
//			}
//			return 0
//		}
//		p := &memo[c][i][j]
//		if *p != -1 { // 已经计算过
//			return *p
//		}
//		// 水平
//		for i2 := i + 1; i2 < n; i2++ {
//			if ms.query(i, j, i2, m) > 0 {
//				res += dfs(c-1, i2, j)
//			}
//		}
//
//		// 垂直
//		for j2 := j + 1; j2 < m; j2++ {
//			if ms.query(i, j, n, j2) > 0 {
//				res += dfs(c-1, i, j2)
//			}
//		}
//		*p = res % mod // 记忆化
//		return *p
//	}
//	return dfs(k-1, 0, 0)
//}

// 翻译成递推
// 自底向上计算 = 递推（动态规划）
// dfs 改成 f 数组
// 递归改成循环（每个参数对应一层循环）
// 递归边界改成 f 数组的初始值
// dfs(c,i,j) = sum(dfs(c-1,i,j2), (j<j2<n) + sum(dfs(c-1,i2,j), (i<i2<m)
// f[c][i][j] = sum(f[c-1][i][j2]) + sum(f[c-1][i2][j])
func ways(pizza []string, k int) int {
	n, m := len(pizza), len(pizza[0])
	const mod = 1_000_000_007
	f := make([][][]int, k) // dfs(c,i,j) 由 c 来控制大小，c 最多为 k-1
	for c := range f {
		f[c] = make([][]int, n)
		for i := range f[c] {
			f[c][i] = make([]int, m)
		}
	}
	ms := NewMatrixSum(pizza)
	for c := 0; c < k; c++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if c == 0 {
					if ms.query(i, j, n, m) > 0 {
						f[c][i][j] = 1
					}
					continue
				}
				res := 0
				// 水平
				for i2 := i + 1; i2 < n; i2++ {
					if ms.query(i, j, i2, m) > 0 {
						res += f[c-1][i2][j]
					}
				}
				// 垂直
				for j2 := j + 1; j2 < m; j2++ {
					if ms.query(i, j, n, j2) > 0 {
						res += f[c-1][i][j2]
					}
				}
				f[c][i][j] = res % mod
			}
		}
	}
	return f[k-1][0][0]
}
