package _60__只出现一次的数字_III

// 思路 1：哈希表（忽略）
// 思路 2：位运算
// 1. 先对所有数字进行一次异或，得到两个出现一次的数字的异或值
// 2. 在异或结果中找到任意为 1 的位
// 3. 根据这一位对所有的数字进行分组
// 4. 在每个组内进行异或操作，得到两个数字

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	// xorSum = a ^ b
	lsb := xorSum & -xorSum // 最低位的 1
	a, b := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
