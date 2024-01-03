package main

// 思路：哈希表
func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]struct{})
	for _, r := range jewels {
		m[r] = struct{}{}
	}
	res := 0
	for _, r := range stones {
		if _, ok := m[r]; ok {
			res++
		}
	}
	return res
}
