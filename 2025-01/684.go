// Runtime : 0 ms
// Memory  : 4.62 MB
package main

func find(p int, s []int) int {
	x := p
	for s[x] != x {
		x = s[x]
	}
	return x
}

func union(p, q int, s []int) bool {
	pRoot, qRoot := find(p, s), find(q, s)
	if pRoot != qRoot {
		s[pRoot] = qRoot
		return true
	}
	return false
}

func findRedundantConnection(edges [][]int) []int {
	s := make([]int, len(edges)+1)
	for p := 1; p < len(edges)+1; p++ {
		s[p] = p
	}

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		if !union(src, dst, s) {
			return []int{src, dst}
		}
	}
	return edges[len(edges)-1]
}
