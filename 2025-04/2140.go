// Runtime : 6 ms
// Memory  : 20.87 MB
package apr25

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n)
	maxPoints := int64(0)

	for i := n - 1; i >= 0; i-- {
		solvedPoint := int64(questions[i][0])
		if next := i + questions[i][1] + 1; next < n {
			solvedPoint += dp[next]
		}
		maxPoints = max(maxPoints, solvedPoint)
		dp[i] = maxPoints
	}

	return maxPoints
}
