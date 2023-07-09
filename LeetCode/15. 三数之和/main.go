package main

import "sort"

// 三数之和
// 思路：排序 + 双指针
// 为了保证不重复：第二重循环枚举的元素要>= 第一重循环枚举的元素
// 为了保证不重复：第三重循环枚举的元素要>= 第二重循环枚举的元素
// 枚举的三元组(a,b,c) 满足 a<=b<=c
// 那么就需要先排序
// 然后双指针降低时间复杂度
// 优化：跳过重复数字

func threesum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { //跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 因为是从小到大排序，说明后面的不可能
			break
		}
		if x+nums[n-1]+nums[n-2] < 0 { // 说明 x 这个数字不可能，可以直接往后枚举了
			continue
		}
		j, k := i+1, n-1 // 双指针
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				// 跳过重复数字
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return ans
}