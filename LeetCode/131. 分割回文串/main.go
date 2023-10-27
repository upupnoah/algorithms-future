package _31__分割回文串

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 输入的视角 （逗号选/不选）
func patition1(s string) (ans [][]string) {
	n := len(s)
	var path []string
	var dfs func(int, int)

	// start 表示当前这段回文字段的开始位置
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		// 不选 i 和 i+1 之间的逗号（i = n-1 时一定要选 -> 到 n-1 必须要截断，后面没字符了...
		if i < n-1 {
			dfs(i+1, start)
		}

		// 选 i 和 i+1 之间的逗号（把s[i] 作为子串的最后一个字符
		if isPalindrome(s[start : i+1]) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)             // 下一个子串从 i+1 开始
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, 0)
	return
}

// 答案的角度（枚举选哪个）
func partition(s string) (ans [][]string) {
	n := len(s)
	var path []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 这里不像 78 题的答案视角是因为需要包含所有的字符，因此需要等到最后一个再添加
			ans = append(ans, append([]string(nil), path...))
			return
		}
		// j 表示子串结束位置
		// 规定一个顺序，从 i 开始，避免重复
		for j := i; j < n; j++ {
			if isPalindrome(s[i : j+1]) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
