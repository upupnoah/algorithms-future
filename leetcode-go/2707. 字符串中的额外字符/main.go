package min_extra_char

// dfs(i): 表示 s 的前 i 个字符的子问题
// 跳过 s 的最后一个字符(把最后一个字符当做额外字符), 有 dfs(i) = dfs(i-1)+1
// 考虑 [枚举选哪个], 如果从s[j] 到 s[i] 的子串在dictionary 中, 有 dfs(i) = min dfs(j-1)  (j = 0,1,...i)
// 递归边界: dfs(-1) = 0
// ans: dfs(n-1)
func minExtraChar(s string, dictionary []string) int {
	has := make(map[string]bool)
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		// 不选第 i 个字符
		res := dfs(i-1) + 1

		// 枚举选哪个
		for j := 0; j <= i; j++ {
			if has[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}
		*p = res
		return res
	}
	return dfs(n - 1)
}
