package _582__Pass_the_Pillow

// 思路：找数学规律
// 答案是在 2 -> n 与 n-1 -> 1, 总共有 2n-2 个可能的值
// 看一下 time % (2n-2) 的值

func passThePillow(n int, time int) int {
	res := time % (2*n - 2)
	if res < n {
		return res + 1
	}
	return 2*n - res - 1 // (n-1-(res-n))
}
