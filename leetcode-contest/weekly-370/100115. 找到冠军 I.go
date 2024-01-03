package main

// func findChampion1(grid [][]int) (ans int) {
// 	cache := make([]int, len(grid))
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid); j++ {
// 			if i != j{
// 				if grid[i][j] == 1 {
// 					cache[j] ++
// 				} else {
// 					cache[i]++
// 				}
// 			}
// 		}
// 	}
// 	for i, v := range cache {
// 		if v == 0 {
// 			return i
// 		}
// 	}
// 	return 
// }