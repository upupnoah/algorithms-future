package _003__每棵子树内缺失的最小基因值

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// todo 启发式合并
func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	var dfs func(int) (map[int]bool, int)
	dfs = func(node int) (map[int]bool, int) {
		geneSet := map[int]bool{nums[node]: true}
		for _, child := range children[node] {
			childGeneSet, y := dfs(child)
			res[node] = max(res[node], y)
			if len(childGeneSet) > len(geneSet) {
				geneSet, childGeneSet = childGeneSet, geneSet
			}
			for gene := range childGeneSet {
				geneSet[gene] = true
			}
		}
		for geneSet[res[node]] {
			res[node]++
		}
		return geneSet, res[node]
	}

	dfs(0)
	return res
}
