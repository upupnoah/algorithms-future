package _06__累加数

import "strconv"

// 输入视角(选/不选)
func isAdditiveNumber1(num string) bool {
	n := len(num)
	var path []int
	var dfs func(int, int) bool
	dfs = func(i, start int) bool {
		if i == n {
			m := len(path)
			if m >= 3 {
				for i := 2; i < m; i++ {
					if path[i] != path[i-1]+path[i-2] {
						return false
					}
				}
				return true
			}
			return false
		}
		if i-start > 0 && num[start] == '0' { // > 2 个字符,并且开头是 0 的数字不合法
			return false
		}
		// 不选 i
		if i < n-1 {
			if dfs(i+1, start) {
				return true
			}
		}
		// 选 i
		digit, _ := strconv.Atoi(num[start : i+1])
		if len(path) < 2 || digit == path[len(path)-1]+path[len(path)-2] {
			path = append(path, digit)
			if dfs(i+1, i+1) {
				return true
			}
			path = path[:len(path)-1]
		}
		return false
	}
	return dfs(0, 0)
}

// 答案视角(枚举选哪个)
func isAdditiveNumber(num string) bool {
	n := len(num)
	var path []int
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			m := len(path)
			if m >= 3 {
				for i := 2; i < m; i++ {
					if path[i] != path[i-1]+path[i-2] {
						return false
					}
				}
				return true
			}
			return false
		}
		m := len(path)
		for j := i; j < n; j++ {
			if j-i > 0 && num[i] == '0' {
				continue
			}
			digit, _ := strconv.Atoi(num[i : j+1])
			if m <= 2 || digit == path[m-1]+path[m-2] {
				path = append(path, digit)
				if dfs(j + 1) {
					return true
				}
				path = path[:len(path)-1]
			}
		}
		return false
	}
	return dfs(0)
}

//输入视角(选/不选)
//func isAdditiveNumber(num string) bool {
//	n := len(num)
//	var path []int
//	var ans [][]int
//	var dfs func(int, int)
//	dfs = func(i, start int) {
//		if i == n {
//			if len(path) > 2 {
//				ans = append(ans, append([]int(nil), path...))
//			}
//			return
//		}
//		// 不选
//		if i < n-1 {
//			dfs(i+1, start)
//		}
//		if i-start > 0 && num[start] == '0' {
//			return
//		}
//		digit, _ := strconv.Atoi(num[start : i+1])
//		// 选
//		if len(path) < 2 || digit == path[len(path)-1]+path[len(path)-2] {
//			path = append(path, digit)
//			dfs(i+1, i+1)
//			path = path[:len(path)-1]
//		}
//
//	}
//	dfs(0, 0)
//	for _, v := range ans {
//		flag := false
//		for i := 2; i < len(v); i++ {
//			if v[i] != v[i-1]+v[i-2] {
//				flag = true
//				break
//			}
//		}
//		if !flag {
//			// fmt.Println(v)
//			return true
//		}
//	}
//	return false
//}

// 答案视角(枚举选哪个)
//func isAdditiveNumber1(num string) bool {
//	n := len(num)
//	var path []int
//	var dfs func(int) bool
//	dfs = func(u int) bool {
//		m := len(path)
//		if m >= 3 && path[m-1] != path[m-2]+path[m-3] { // 剪枝
//			return false
//		}
//		if u >= n {
//			return m >= 3
//		}
//
//		var k int
//		if num[u] == '0' {
//			k = u
//		} else {
//			k = n - 1
//		}
//		for i := u; i <= k; i++ {
//			p, _ := strconv.Atoi(num[u : i+1]) // Current position
//			path = append(path, p)
//			if dfs(i + 1) {
//				return true
//			}
//			path = path[:len(path)-1]
//		}
//		return false
//	}
//	return dfs(0)
//}

// 答案视角(枚举选哪个)
//func isAdditiveNumber2(num string) bool {
//	n := len(num)
//	var path []int
//	var ans [][]int
//	var dfs func(int)
//	dfs = func(i int) {
//		if i == n {
//			if len(path) > 2 {
//				ans = append(ans, append([]int(nil), path...))
//			}
//		}
//		m := len(path)
//		for j := i; j < n; j++ {
//			if j-i > 0 && num[i] == '0' {
//				continue
//			}
//			digit, _ := strconv.Atoi(num[i : j+1])
//			if m <= 2 || digit == path[m-1]+path[m-2] {
//				path = append(path, digit)
//				dfs(j + 1)
//				path = path[:len(path)-1]
//			}
//		}
//	}
//	dfs(0)
//	for _, v := range ans {
//		flag := false
//		for i := 2; i < len(v); i++ {
//			if v[i] != v[i-1]+v[i-2] {
//				flag = true
//				break
//			}
//		}
//		if !flag {
//			return true
//		}
//	}
//	return false
//}
