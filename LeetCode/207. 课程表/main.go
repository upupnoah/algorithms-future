package _07__课程表

// 思路：拓扑排序
// 判断最后的序列长度是否等于课程数
// BFS 的应用
// 1. 将所有入度为 0 的节点加入队列
// 2. 将队列中的节点一个一个的释放，释放的同时将该节点指向的节点的入度减一
// 3. 如果入度减为 0，则将该节点加入队列
// 4. 重复 2，3 步骤，直到队列为空
// 5. 判断释放的节点数是否等于课程数

// 使用邻接表存稀疏图

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	d := make([]int, numCourses) // 入度
	for _, edge := range prerequisites {
		a, b := edge[0], edge[1]
		edges[b] = append(edges[b], a)
		d[a]++
	}
	var q []int
	for i, v := range d {
		if v == 0 {
			q = append(q, i)
		}
	}
	var res []int
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		res = append(res, t)
		for _, p := range edges[t] {
			d[p]--
			if d[p] == 0 {
				q = append(q, p)
			}
		}
	}
	return len(res) == numCourses
}
