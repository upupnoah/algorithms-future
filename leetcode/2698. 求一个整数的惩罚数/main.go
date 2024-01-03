package _698__求一个整数的惩罚数

import "strconv"

var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int) bool
		sum := 0
		// 答案视角,枚举选哪个
		// 返回值是 bool 可以用来剪枝
		dfs = func(p int) bool {
			if p == n {
				return sum == i
			}
			// 枚举选哪个
			x := 0
			for j := p; j < n; j++ {
				x = x*10 + int(s[j]&15)
				sum += x
				if dfs(j + 1) {
					return true
				}
				sum -= x
			}
			return false
		}

		preSum[i] = preSum[i-1]
		if dfs(0) {
			preSum[i] += i * i // 计算前缀和
		}

		// 输入视角(选/不选)
		//var dfs func(int, int, int) bool
		//dfs = func(p, start, sum int) bool {
		//	if p == n {
		//		return sum == i
		//	}
		//	// 不选
		//	if p < n-1 {
		//		if dfs(p+1, start, sum) {
		//			return true
		//		}
		//	}
		//
		//	// 选
		//	x := 0
		//	for j := start; j <= p; j++ {
		//		x = x*10 + int(s[j]&15)
		//	}
		//	if dfs(p+1, p+1, sum+x) { // sum 在参数里面,每次都会拷贝一份,因此不需要恢复现场
		//		return true
		//	}
		//	return false
		//}
		// 前缀和
		//preSum[i] = preSum[i-1]
		//if dfs(0, 0, 0) {
		//	preSum[i] += i * i // 计算前缀和
		//}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
