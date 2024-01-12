package main

// 哈希表 hash
func findRepeatedDnaSequences(s string) (ans []string) {
	hash := make(map[string]int)
	for i := 0; i+9 < len(s); i++ {
		hash[s[i:i+10]]++
	}
	for k, v := range hash {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return
}
