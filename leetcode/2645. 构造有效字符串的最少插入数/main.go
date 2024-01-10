package add_minimum

// 方法一: 考虑 abc 的个数
// 答案由 t 个 ‘abc’ 组成, 因此需要插入的字符个数最少是 3*t - len(s)
// 如果s[i-1] >= s[i], 则需要生成一个新的 abc
func addMinimum(word string) int {
	t := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] >= word[i] {
			t++ // 必须生成一个新的 abc
		}
	}
	return t*3 - len(word)
}

// 方法二: 考虑相邻两个字母
// ab  b-a-1 = 0 个 (不需要插入)
// ac  c-a-1 = 1 个 (需要插入一个 b)
// ca  a-c-1 (+3) = 0, 但是我们需要的 0(ca 之间不需要插入), 因此 + 3
// 综上: (右-左+2) % 3
// 坑点: 开头 和 结尾是没有相邻两个字符的, 因此需要特殊判断
// s[0]: s[0]-'a'
// s[n-1]: 'c'-s[n-1]
// s[0] + s[n-1] = s[0]-s[n-1]+2
func addMinimum1(word string) int {
	ans := int(word[0]) - int(word[len(word)-1]) + 2
	for i := 1; i < len(word); i++ {
		ans += (int(word[i]) - int(word[i-1]) + 2) % 3
	}
	return ans
}
