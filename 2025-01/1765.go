// Runtime : 48 ms
// Memory  : 51.20 MB

package main

type point struct {
	i int
	j int
}

func highestPeak(isWater [][]int) [][]int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	mHeight, mWidth := len(isWater), len(isWater[0])
	queue := []point{}

	heightMap := make([][]int, mHeight)
	for i := range heightMap {
		heightMap[i] = make([]int, mWidth)
		for j := range heightMap[i] {
			if isWater[i][j] == 1 {
				heightMap[i][j] = 0
				queue = append(queue, point{i, j})
			} else {
				heightMap[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		row, col := front.i, front.j

		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < mHeight &&
				newCol >= 0 && newCol < mWidth &&
				heightMap[newRow][newCol] == -1 {

				heightMap[newRow][newCol] = heightMap[row][col] + 1
				queue = append(queue, point{newRow, newCol})
			}
		}
	}

	return heightMap
}
