package main

import (
	"fmt"
	"math"
)

func minimumSum1(nums []int) int {
	n := len(nums)
	res := math.MaxInt32
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					res = min(res, nums[i]+nums[j]+nums[k])
				}
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func main() {
	//输入：nums = [8,6,1,5,3]
	//输出：9
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))

	//输入：nums = [5,4,8,7,10,2]
	//输出：13
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))

}
