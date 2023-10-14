package _37__Single_Number_II

// 思路：统计每一位上1的个数，如果是3的倍数，说明目标值在该位上是0，否则是1
// 由于题目限定了32位，所以可以用int32来存储

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
