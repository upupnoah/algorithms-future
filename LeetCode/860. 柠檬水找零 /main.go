package main

// 思路：贪心
// 策略：对于 20，优先找零 10 元，再找零 5 元
//
//	对于 10， 找零 5 元
//
// 做法：统计 5 和 10 的数量
func lemonadeChange(bills []int) bool {
	cnt5, cnt10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			cnt5++
		} else if bill == 10 {
			if cnt5 == 0 {
				return false
			}
			cnt5--
			cnt10++
		} else {
			if cnt10 > 0 && cnt5 > 0 {
				cnt10--
				cnt5--
			} else if cnt5 >= 3 {
				cnt5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
