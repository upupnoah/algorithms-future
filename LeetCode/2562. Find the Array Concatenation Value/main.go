package _562__Find_the_Array_Concatenation_Value

// 我的写法

func findTheArrayConcVal(nums []int) int64 {
	res := 0
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t := nums[j]
		x := 1
		for t != 0 {
			t /= 10
			x *= 10
		}
		res += nums[i]*x + nums[j]
	}
	if n%2 == 1 {
		res += nums[n/2]
	}
	return int64(res)
}

// 0x3f 的

func findTheArrayConcVal2(nums []int) (res int64) {
	for len(nums) > 1 {
		x := nums[0]
		for y := nums[len(nums)-1]; y != 0; y /= 10 {
			x *= 10
		}
		res += int64(x + nums[len(nums)-1])
		nums = nums[1 : len(nums)-1]
	}
	if len(nums) == 1 { // 奇数情况还剩下最中间的元素
		res += int64(nums[0])
	}
	return
}
