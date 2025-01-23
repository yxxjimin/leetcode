// Runtime : 0 ms
// Memory  : 9.13 MB
package main

func isColumnConnected(grid [][]int, col int) bool {
	colCnt := 0
	for i := range grid {
		colCnt += grid[i][col]
	}
	return colCnt > 1
}

func countServers(grid [][]int) int {
	totalCnt := 0
	checkCol := make([]int, 0)

	for _, row := range grid {
		lastCol := 0
		rowCnt := 0
		for j := range row {
			if row[j] == 1 {
				rowCnt += 1
				lastCol = j
			}
		}
		if rowCnt == 1 {
			checkCol = append(checkCol, lastCol)
		}
		totalCnt += rowCnt
	}

	for _, col := range checkCol {
		if !isColumnConnected(grid, col) {
			totalCnt -= 1
		}
	}

	return totalCnt
}
