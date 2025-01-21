package main

import "math"

// Runtime : 23 ms
// Memory  : 18.86 MB
func gridGame(grid [][]int) int64 {
	length := len(grid[0])
	sScore := int64(math.MaxInt64)

	lbPrefixSum := make([]int64, length+1)
	rtPrefixSum := make([]int64, length+1)

	for i := 0; i < length; i++ {
		lbPrefixSum[i+1] = lbPrefixSum[i] + int64(grid[1][i])
		rtPrefixSum[i+1] = rtPrefixSum[i] + int64(grid[0][i])
	}

	for pvt := 0; pvt < length; pvt++ {
		lbSum := lbPrefixSum[pvt]
		rtSum := rtPrefixSum[length] - rtPrefixSum[pvt+1]

		sChoice := max(lbSum, rtSum)
		sScore = min(sScore, sChoice)
	}

	return sScore
}

// Runtime : 0 ms
// Memory  : 9.46 MB
// Comment :
//   - 우상단 총합을 처음부터 구해놓고 하나 씩 빼면 prefixSum 슬라이스를 만들 필요가 없어짐
//   - 좌하단 총합은 그냥 0부터 차례대로 누적시키면 됨
func gridGame2(grid [][]int) int64 {
	rtSum, lbSum, sScore := int64(0), int64(0), int64(math.MaxInt64)

	for _, num := range grid[0] {
		rtSum += int64(num)
	}

	for pvt := 0; pvt < len(grid[0]); pvt++ {
		rtSum -= int64(grid[0][pvt])
		sScore = min(sScore, max(rtSum, lbSum))
		lbSum += int64(grid[1][pvt])
	}

	return sScore
}
