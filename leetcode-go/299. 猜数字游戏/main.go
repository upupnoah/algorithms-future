package main

import (
	"fmt"
)

// https://leetcode.cn/problems/bulls-and-cows/description/?envType=daily-question&envId=2024-03-10
// 直接统计多少位是相等的
// 不相等的用两个 hash map 来存, 最后 cow 的数量为 sigma(min(cnt1[i], cnt2[i]))
func getHint(secret, guess string) string {
    a, b := 0, 0
    var cntS, cntG [10]int
    for i := range secret {
        if secret[i] == guess[i] {
            a++
        } else {
            cntS[secret[i]-'0']++
            cntG[guess[i]-'0']++
        }
    }
    for i := 0; i < 10; i++ {
        b += min(cntS[i], cntG[i])
    }
    return fmt.Sprintf("%dA%dB", a, b)
}