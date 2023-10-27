package _601__最多可达成的换楼请求数目

// 使用 cache 数组，进来的和出去的人一样，意味着 cache数组的值都为 0
// 出去-1， 进来+1

func maximumRequests1(n int, requests [][]int) (ans int) {
	cache := make([]int, n)
	var dfs func(int, int)
	dfs = func(i, cnt int) {
		if i == len(requests) {
			for _, v := range cache {
				if v != 0 {
					return
				}
			}
			ans = max(ans, cnt)
			return
		}
		// 选第 i 个request
		cache[requests[i][0]]--
		cache[requests[i][1]]++
		dfs(i+1, cnt+1)
		cache[requests[i][0]]++
		cache[requests[i][1]]--

		// 不选
		dfs(i+1, cnt)
	}
	dfs(0, 0)
	return
}

func maximumRequests(n int, requests [][]int) (ans int) {
	cache := make([]int, n)
	var dfs func(int, int)
	//check := func() bool {
	//	for _, v := range cache {
	//		if v != 0 {
	//			return false
	//		}
	//	}
	//	return true
	//}
	dfs = func(i, cnt int) {
		//if check() {
		//	ans = max(ans, cnt)
		//}
		flag := false
		for _, v := range cache {
			if v != 0 {
				flag = true
				break
			}
		}
		if !flag {
			ans = max(ans, cnt)
		}
		if i == len(requests) {
			return
		}
		for j := i; j < len(requests); j++ {
			cache[requests[j][0]]--
			cache[requests[j][1]]++
			dfs(j+1, cnt+1)
			cache[requests[j][1]]--
			cache[requests[j][0]]++
		}
	}
	dfs(0, 0)
	return
}
