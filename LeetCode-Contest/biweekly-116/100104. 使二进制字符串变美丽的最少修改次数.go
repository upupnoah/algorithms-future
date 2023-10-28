package biweekly_116

// 贪心:遇到两个一样的就切,不一样ans++
// 步长为 2
func minChanges(s string) int {
	ans := 0
	for i := 0; i < len(s)-1; {
		if s[i] != s[i+1] {
			ans++
		}
		i = i + 2
	}
	return ans
}
