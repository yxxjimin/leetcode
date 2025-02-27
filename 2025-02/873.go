// Runtime : 108 ms
// Memory  : 33.55 MB
package feb25

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	maxLength := 0

	for k := 2; k < n; k++ {
		i, j := 0, k-1
		for i < j {
			if sum := arr[i] + arr[j]; sum < arr[k] {
				i++
			} else if sum > arr[k] {
				j--
			} else {
				dp[j][k] = dp[i][j] + 1
				maxLength = max(maxLength, dp[j][k])
				i++
				j--
			}
		}
	}

	if maxLength > 0 {
		return maxLength + 2
	}
	return 0
}
