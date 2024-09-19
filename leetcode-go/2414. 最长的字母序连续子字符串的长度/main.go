package main

// 思路:
// 很直观就能想到(连续)子序列问题, 用双指针/滑动窗口
// 遍历字符串, 维护 cnt 和 ans


func longestContinuousSubstring(s string) int {
	ans, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i-1]+1 == s[i] {
			cnt += 1
			ans = max(ans, cnt)
		} else {
			cnt = 1
		}
	}
	return ans
}
