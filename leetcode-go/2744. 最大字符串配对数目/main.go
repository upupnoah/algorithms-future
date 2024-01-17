package maximum_number_of_string_pairs

func maximumNumberOfStringPairs(words []string) int {
	st := [26][26]bool{}
	ans := 0
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		// 因为字符串互不相同, 所以能够匹配一定是反转的 -> "ab" "ba"
		if st[y][x] {
			ans++
		} else {
			st[x][y] = true
		}
	}
	return ans
}
