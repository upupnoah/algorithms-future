package _678__老人的数目

func countSeniors(details []string) int {
	ans := 0
	for _, s := range details {
		if s[12]-'0' >= 1 && s[11]-'0' >= 6 || s[11]-'0' > 6 {
			ans++
		}
	}
	return ans
}

func countSeniors2(details []string) (ans int) {
	for _, s := range details {
		// 对于数字字符，&15 等价于 -'0'，但是不需要加括号
		if s[11]&15*10+s[12]&15 > 60 {
			ans++
		}
	}
	return
}
