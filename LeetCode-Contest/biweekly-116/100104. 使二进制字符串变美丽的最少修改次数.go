package biweekly_116

// 贪心:遇到两个一样的就切,不一样ans++
// 步长为 2
func minChanges(s string) (ans int) {
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			ans++
		}
	}
	return
}

// 思考题：在修改次数最少的前提下，最少可以分割成多少段（每段字符都一样）
func minSegment(s string) (ans int) {
	var st byte = '#'
	n := len(s) - 1
	if s[0] == s[1] {
		st = s[0]
		ans++
	} else {
		if 2 < n {
			st = s[2]
			ans++
		} else {
			return 1
		}
	}
	for i := 2; i < n; i += 2 {
		if s[i] == s[i+1] && s[i] != st {
			ans++
			st = s[i]
		}
		// 如果修改,则必然可以与前面的合成一段
		// s[i] != s[i+1] && (st == s[i] or s[i+1)) -> continue
		// s[i] == s[i+1] && s[i] == st -> continue

	}
	return
}
