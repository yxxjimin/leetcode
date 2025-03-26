// Runtime : 37 ms
// Memory  : 20.96 MB
package mar25

import (
	"math"
	"slices"
)

func minOperations2(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	flat := make([]int, m*n)
	offset := grid[0][0] % x

	for i := range grid {
		row := grid[i]
		for j, num := range row {
			if offset != num%x {
				return -1
			}
			flat[i*n+j] = num / x
		}
	}

	slices.Sort(flat)
	median := flat[len(flat)/2]
	ans := float64(0)

	for _, num := range flat {
		ans += math.Abs(float64(num) - float64(median))
	}

	return int(ans)
}
