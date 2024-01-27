package can_measure_water

// DFS
func canMeasureWater(x int, y int, t int) bool {
	var dfs func(int, int) bool
	memo := make(map[[2]int]bool)
	dfs = func(remainX, remainY int) bool {
		if remainX+remainY == t {
			return true
		}
		if memo[[2]int{remainX, remainY}] {
			return false
		}
		memo[[2]int{remainX, remainY}] = true
		// 尝试所有操作
		// 倒满任意一个容器
		if dfs(x, remainY) || dfs(remainX, y) {
			return true
		}
		// 倒空任意一个容器
		if dfs(0, remainY) || dfs(remainX, 0) {
			return true
		}
		// 将 x 倒入 y
		if dfs(max(remainX-(y-remainY), 0), min(remainX+remainY, y)) {
			return true
		}
		// 将 y 倒入 x
		if dfs(min(x, remainX+remainY), max(0, remainY-(x-remainX))) {
			return true
		}
		return false
	}
	return dfs(0, 0)
}

// 贝祖定理: ax + by = z 有解, 则 z 是 x, y 的最大公约数的倍数(充要条件)
func canMeasureWater1(x int, y int, t int) bool {
	if x+y < t {
		return false
	}
	// 贝祖定理
	// ax + by = t 有解当且仅当 t 是 x, y 的最大公约数的倍数
	return t%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
