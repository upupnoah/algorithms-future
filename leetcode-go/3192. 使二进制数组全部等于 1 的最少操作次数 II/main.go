package main

func minOperations(nums []int) (k int) {
	for _, x := range nums {
		if x == k%2 { // 必须操作
			k++
		}
	}
	return
}
