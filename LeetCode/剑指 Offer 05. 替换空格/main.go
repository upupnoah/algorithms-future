package main

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func replaceSpace(s string) string {
	var ans string
	for _, v := range s {
		if v == ' ' {
			ans += "%20"
		} else {
			ans += string(v)
		}
	}
	return ans	
}