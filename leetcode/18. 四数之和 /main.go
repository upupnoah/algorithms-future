package main

import "sort"

// 思路：i从 0 开始，j 从 1 开始遍历
// k 从 j+1， l 从 n-1 开始 （双指针）
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for a := 0; a < n-3; a++ {
		x := nums[a]
		if a > 0 && x == nums[a-1] {
			continue
		}
		if x+nums[a+1]+nums[a+2]+nums[a+3] > target {
			break
		}
		if x+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			y := nums[b]
			if b > a+1 && y == nums[b-1] {
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target {
				break
			}
			if x+y+nums[n-1]+nums[n-2] < target {
				continue
			}
			l, r := b+1, n-1
			for l < r {
				sum := x + y + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					ans = append(ans, []int{x, y, nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; r > l && nums[r] == nums[r+1]; r-- {
					}
				}
			}
		}
	}
	return ans
}
