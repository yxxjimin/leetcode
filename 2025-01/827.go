// Runtime : 101 ms
// Memory  : 16.16 MB
package main

type DisjointSet struct {
	prnt []int
	size []int
}

func NewDisjointSet(n int) *DisjointSet {
	prnt, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		prnt[i] = i
		size[i] = 1
	}
	return &DisjointSet{prnt, size}
}

func (ds *DisjointSet) Find(x int) int {
	if x != ds.prnt[x] {
		ds.prnt[x] = ds.Find(ds.prnt[x])
	}
	return ds.prnt[x]
}

func (ds *DisjointSet) Union(x, y int) {
	rootX, rootY := ds.Find(x), ds.Find(y)
	if rootX == rootY {
		return
	}
	if ds.size[rootX] > ds.size[rootY] {
		ds.prnt[rootY] = rootX
		ds.size[rootX] += ds.size[rootY]
	} else {
		ds.prnt[rootX] = rootY
		ds.size[rootY] += ds.size[rootX]
	}
}

func isIsland(r, c int, grid [][]int) bool {
	n := len(grid)
	return 0 <= r && r < n &&
		0 <= c && c < n &&
		grid[r][c] == 1
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	ds := NewDisjointSet(n * n)
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				node := r*n + c
				for _, dir := range dirs {
					nr, nc := r+dir[0], c+dir[1]
					if isIsland(nr, nc, grid) {
						neighbor := nr*n + nc
						ds.Union(node, neighbor)
					}
				}
			}
		}
	}

	maxSize := -1
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				neighbors := make(map[int]bool)
				for _, dir := range dirs {
					nr, nc := r+dir[0], c+dir[1]
					if isIsland(nr, nc, grid) {
						root := ds.Find(nr*n + nc)
						neighbors[root] = true
					}
				}

				currSize := 1
				for root := range neighbors {
					currSize += ds.size[root]
				}
				maxSize = max(maxSize, currSize)
			}
		}
	}

	if maxSize == -1 {
		return n * n
	}
	return maxSize
}
