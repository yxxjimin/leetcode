// Runtime : 82 ms
// Memory  : 11.74 MB
package main

func numGrpsAsRoot(src, n int, graph [][]int) int {
	q, d, depth := []int{src}, 1, make([]int, n+1)
	depth[src] = d

	for len(q) > 0 {
		lq := len(q)
		for i := 0; i < lq; i++ {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if depth[v] == 0 {
					q = append(q, v)
					depth[v] = d + 1
				} else if depth[v] == depth[u] {
					return -1
				}
			}
		}
		d++
	}
	return d - 1
}

func find2(p int, parent []int) int {
	for p != parent[p] {
		p = parent[p]
	}
	return p
}

func union2(p, q int, parent []int) {
	pRoot, qRoot := find2(p, parent), find2(q, parent)
	if pRoot != qRoot {
		parent[qRoot] = pRoot
	}
}

func magnificentSets(n int, edges [][]int) int {
	graph := make([][]int, n+1)
	parent := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
		union2(src, dst, parent)
	}

	maxGrpOfComponent := make(map[int]int)
	for u := 1; u < n+1; u++ {
		if grpsCnt := numGrpsAsRoot(u, n, graph); grpsCnt == -1 {
			return -1
		} else {
			uRoot := find2(u, parent)
			maxGrpOfComponent[uRoot] = max(maxGrpOfComponent[uRoot], grpsCnt)
		}
	}

	ans := 0
	for _, v := range maxGrpOfComponent {
		ans += v
	}

	return ans
}
