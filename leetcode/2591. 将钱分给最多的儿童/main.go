package _591__将钱分给最多的儿童

// 思路1：二分答案 [0, children]
// 思路2：贪心
func distMoney(money int, children int) int {
	//l, r := 0, children
	//for l < r {
	//	mid := l + (r-l+1)>>1
	//	t := money - mid*8
	//	remain := children - mid
	//	if mid*8 > money || t < remain || (remain == 1 && t == 4) {
	//		r = mid - 1
	//	} else {
	//		l = mid
	//	}
	//}
	//return l

	// 贪心写法
	money -= children
	if money < 0 {
		return -1
	}
	ans := min(money/7, children) // 给每个人分配 7 美元
	money -= ans * 7
	children -= ans // 剩余的人
	if children == 0 && money > 0 || children == 1 && money == 3 {
		ans--
	}
	return ans
}
