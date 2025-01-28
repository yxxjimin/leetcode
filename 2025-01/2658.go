// Runtime : 0 ms
// Memory  : 8.29 MB
package main

type point2 struct {
	r, c int
}

func hasFish(r, c int, grid [][]int) bool {
	return 0 <= r && r < len(grid) &&
		0 <= c && c < len(grid[0]) &&
		grid[r][c] > 0
}

func dfs(src point2, grid [][]int) int {
	dirs := []point2{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	fish, stack := 0, []point2{src}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if grid[curr.r][curr.c] != 0 {
			fish += grid[curr.r][curr.c]
			grid[curr.r][curr.c] = 0
			for _, dir := range dirs {
				nr, nc := curr.r+dir.r, curr.c+dir.c
				if hasFish(nr, nc, grid) {
					stack = append(stack, point2{nr, nc})
				}
			}
		}
	}
	return fish
}

func findMaxFish(grid [][]int) int {
	maxFish := 0
	for r, row := range grid {
		for c := range row {
			if grid[r][c] > 0 {
				maxFish = max(maxFish, dfs(point2{r, c}, grid))
			}
		}
	}
	return maxFish
}
