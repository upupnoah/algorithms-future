package _520__统计能整除数字的位数

// 简单模拟
// %10 /10

func countDigits(num int) int {
	ans := 0
	t := num
	for num != 0 {
		if t%(num%10) == 0 {
			ans++
		}
		num /= 10
	}
	return ans
}
