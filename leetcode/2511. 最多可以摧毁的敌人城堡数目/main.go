package _511__最多可以摧毁的敌人城堡数目

// 我的思路：
// 因为只能从我控制的城堡开始走，因此对于所有我控制的城堡，往走，往右，与 res 取一个最大值
func captureForts(forts []int) int {
	var res int
	n := len(forts)
	for i := 0; i < n; i++ {
		if forts[i] == 1 {
			j, k := i-1, i+1
			cnt := 0
			for j >= 0 && forts[j] == 0 {
				cnt++
				j--
			}
			if j >= 0 && forts[j] == -1 {
				res = max(res, cnt)
			}
			cnt = 0
			for k < n && forts[k] == 0 {
				cnt++
				k++
			}
			if k < n && forts[k] == -1 {
				res = max(res, cnt)
			}
		}
	}
	return res
}

// 更好的思路：一次遍历，我们只需要知道 i 和 j 之间的距离是多少
// 只需要知道(−1,1)或者(1,−1)之间0的距离，那就是两个端点为−1或1
// 记录每次遍历到 1 或者 -1的位置 pre，forts[pre] != forts[i] -> res = max(res, i-pre-1)
func captureForts2(forts []int) int {
	res, pre := 0, -1
	n := len(forts)
	for i := 0; i < n; i++ {
		if forts[i] != 0 {
			if pre != -1 && forts[i] != forts[pre] {
				res = max(res, i-pre-1)
			}
			pre = i
		}
	}
	return res
}
