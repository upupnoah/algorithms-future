package main

import (
	"math"
	"sort"
)

// 思路：排序 + 双指针
// 如何判断a，b 哪个更接近 target？|a-target| < |b-target|，则 a 更接近
// 每次 指针向右或者向左移动的时候，如果状态改变，就更新一下答案(太麻烦)
// 直接：对于每个 sum，都判断是否需要更新 res
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt32
	for i, x := range nums[:n-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := x + nums[l] + nums[r]
			if sum == target{
				return target
			}
			// 对于每个 sum，更新 res
			if abs(res-target) > abs(sum-target) {
				res = sum
			}
			// 优化
			// if sum < target {
			// 	l++
			// }
			// if sum > target {
			// 	r--
			// }
			// 移动到下一个不相等的位置
			if sum < target {
				l++
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}
			if sum > target {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}

	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	}
	return x
}