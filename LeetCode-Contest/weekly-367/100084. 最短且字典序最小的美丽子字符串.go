package main

import "strings"

// 方法 1：枚举
// 首先，如果 s 中 1 的个数不足 k，直接返回空串
// 否则一定有解
// 从 k 开始枚举答案的长度 size，然后在 s 中枚举所有长为 size 的子串
// 同时维护字典序最小的子串。如果存在一个子串，其中 1  的个数等于 k，则返回字典序最小的子串

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	for size := k; ; size++ {
		ans := ""
		for i := 0; i+size <= len(s); i++ {
			t := s[i : i+size]
			if (ans == "" || t < ans) && strings.Count(t, "1") == k {
				ans = t
			}
		}
		if ans != "" {
			return ans
		}
	}
}

// 方法二：滑动窗口
// 由于答案中恰好有 k 个 1，我们也可以用滑动窗口找最短的答案
// 如果窗口内的 1 的个数超过 k，或者窗口端点是 0，就可以缩小窗口
func shortestBeautifulSubstring1(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s // 初始化为最大的子串（s 本身）
	cnt := 0
	left := 0
	for right, b := range s {
		cnt += int(b & 1)
		for cnt > k || s[left] == '0' {
			cnt -= int(s[left] & 1)
			left++
		}
		if cnt == k {
			t := s[left : right+1]
			// t < ans 这部分时间爱你复杂度是 O(N) 的，可以用字符串哈希 || 后缀数组优化成 O(nlogn)
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}
