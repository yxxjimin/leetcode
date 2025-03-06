// Runtime : 0 ms
// Memory  : 7.14 MB
package mar25

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	found, ans := make([]bool, n*n+1), make([]int, 2)
	found[0] = true

	for i := range grid {
		for j := range grid[i] {
			if tgt := grid[i][j]; !found[tgt] {
				found[tgt] = true
			} else {
				ans[0] = tgt
			}
		}
	}

	for i := range found {
		if !found[i] {
			ans[1] = i
			return ans
		}
	}
	return ans
}
