package _11__有效三角形的个数

import "sort"

// 思路： 可以重复
// 公式： 两边之和 > 第三边
// 排序之后：a < b < c
// a + b > c （只需要考虑这个条件）
// a + c > b (一定成立）
// b + c > a （一定成立）

func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 2; i < n; i++ {
		a, b := 0, i-1
		c := nums[i]
		if c > nums[b]+nums[b-1] {
			continue
		}
		for a < b {
			if nums[a]+nums[b] > c {
				ans += b - a
				b--
			} else {
				a++
			}
		}
	}
	return ans
}
