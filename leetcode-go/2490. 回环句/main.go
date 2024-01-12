package main

// 简单判断
func isCircularSentence(sentence string) bool {
	n := len(sentence)
	if sentence[0] != sentence[n-1] {
		return false
	}
	for i := 0; i < n; i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
