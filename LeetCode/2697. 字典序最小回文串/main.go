package make_smallest_palindrome

// 思路: 字典序最小 -> 转换的时候优先换成字典序的(双指针)
func makeSmallestPalindrome(s string) string {
	n := len(s)
	str := []byte(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			if str[i] < str[j] {
				str[j] = str[i]
			} else {
				str[i] = str[j]
			}
		}
	}
	return string(str)
}