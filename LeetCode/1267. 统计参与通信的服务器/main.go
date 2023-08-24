package main

// 思路：两重循环 + 哈希
// 通过 sr，sc 两个数组分别标记每行和每列的服务器数量
// 然后再编译一遍整个矩阵，if grid[i][j] == 1 && sr[i]>1 || sc[j] > 1  ans++
// 时间复杂度：O(n*m)
// 空间复杂度：O(n+m)
func countServers(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	sr, sc := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				sr[i], sc[j] = sr[i]+1, sc[j]+1
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && (sr[i] > 1 || sc[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
