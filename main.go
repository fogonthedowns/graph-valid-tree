package main

func validTree(n int, edges [][]int) bool {
	m := make(map[int][]int)
	vis := make([]bool, n)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}

	// keep track of visited
	if !validTreeHelper(m, vis, 0, -1) {
		return false
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			return false
		}
	}

	return true
}

func validTreeHelper(m map[int][]int, vis []bool, cur int, prev int) bool {
	// if we already visited it, return false
	if vis[cur] {
		return false
	}

	vis[cur] = true
	// range over the neighbors
	for _, node := range m[cur] {
		// skip simple parent/child cycles. a->b b->a
		if node == prev {
			continue
		}

		// return false if its not valid
		if !validTreeHelper(m, vis, node, cur) {
			return false
		}
	}
	return true
}
