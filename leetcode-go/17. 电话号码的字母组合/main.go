package _7__电话号码的字母组合

var mapping = [][]byte{
	{},
	{},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

// 原问题： 构造长为 n 的字符串
//
//	枚举一个字母
//
// 子问题： 构造长为 n-1 的字符串
// 回溯有一个增量构造答案的过程，这个过程通常用递归实现
func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	var path []byte
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		for _, b := range mapping[digits[i]] {
			path[i] = b
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
