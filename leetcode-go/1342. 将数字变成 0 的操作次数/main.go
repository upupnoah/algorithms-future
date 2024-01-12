package main

// 题意：给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。
//
//	如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1
//
// 做法 1：简单模拟
// 做法 2：位运算（本质上就是右移 + 消减 1），观察最后一位是否为 1
// solution1
func numberOfSteps(num int) int {
	cnt := 0
	for num != 0 {
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
		cnt++
	}
	return cnt
}

func numberOfSteps2(num int) int {
	// -1 是因为右移 getLoc 次，最后一次不需要右移
	// ans = 右移次数+消减 1 的次数
	return max(getLoc(num)+getCnt(num)-1, 0)

}

func getLoc(x int) int {
	for i := 31; i >= 0; i-- {
		if ((x >> i) & 1) == 1 {
			return i + 1
		}
	}
	return -1 //never
}

func getCnt(x int) int {
	ans := 0
	for i := 31; i >= 0; i-- {
		if ((x >> i) & 1) == 1 {
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// class Solution {
//     public int numberOfSteps(int num) {
//         return Math.max(getLoc(num) + getCnt(num) - 1, 0);
//     }
//     int getLoc(int x) {
//         for (int i = 31; i >= 0; i--) {
//             if (((x >> i) & 1) == 1) return i + 1;
//         }
//         return -1; // never
//     }
//     int getCnt(int x) {
//         int ans = 0;
//         for (int i = 31; i >= 0; i--) {
//             if (((x >> i) & 1) == 1) ans++;
//         }
//         return ans;
//     }
// }
