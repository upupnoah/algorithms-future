package main

import "strings"

// 脑筋急转弯：
// LR 不能互相穿透对方，因此去掉'_'后应该是一样的，不一样的直接 false
// 双指针
// 从左到右，找到start[i] 和 target[j] （start[i] == target[j]）
// if start [i] == 'L' && i < j  : false
// if start [i] == 'R' && i > j  : false
func canChange(start string, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			//           ( bool    ==  bool  )
			// i != j 等价于 i > j || i < j
			if i != j && (c == 'L' == (i < j)) {
				return false
			}
			j++
		}
	}
	return true
}
