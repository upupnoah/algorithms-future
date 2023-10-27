package _84__字母大小写全排列

// 选/不选， 输入视角
func letterCasePermutation1(s string) (ans []string) {
	n := len(s)
	path := make([]byte, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(append([]byte(nil), path...)))
			return
		}
		if '0' <= s[i] && s[i] <= '9' {
			path[i] = s[i]
			dfs(i + 1)
		} else {
			//if 'a' <= s[i] && s[i] <= 'z' {
			//	path[i] = s[i]
			//	dfs(i + 1)
			//
			//	path[i] = s[i] - 32
			//	dfs(i + 1)
			//} else {
			//	path[i] = s[i]
			//	dfs(i + 1)
			//
			//	path[i] = s[i] + 32
			//	dfs(i + 1)
			//}
			path[i] = s[i]
			dfs(i + 1)

			path[i] ^= 32 // 大写转为小写，小写转为大写
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

// 枚举选哪个，答案视角
func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	path := []byte(s)
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, string(append([]byte(nil), path...)))
		for j := i; j < n; j++ {
			if path[j] >= 'A' {
				path[j] ^= 32
				dfs(j + 1)
				path[j] ^= 32
			}
		}
	}
	dfs(0)
	return
}
