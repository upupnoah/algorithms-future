package main

import (
	"fmt"
	"math"
)

// 前后缀分解

func minimumSum(a []int) (ans int) {
	n := len(a)
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], a[i])
	}
	pre := a[0] // 在遍历的过程中维护一个 前缀最小值
	ans = math.MaxInt32
	for j := 1; j < n-1; j++ {
		if pre < a[j] && suf[j+1] < a[j] {
			ans = min(ans, pre+a[j]+suf[j+1])
		}
		pre = min(pre, a[j])
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return
}

func main() {
	//输入：nums = [8,6,1,5,3]
	//输出：9
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))

	//输入：nums = [5,4,8,7,10,2]
	//输出：13
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))

}
