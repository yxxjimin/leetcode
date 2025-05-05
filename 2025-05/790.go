// Runtime : 0 ms
// Memory  : 4.08 MB
package may25

func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	mod := 1000000007
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	sums := make([]int, n+1)
	sums[0], sums[1], sums[2] = 1, 2, 4

	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*sums[i-3]) % mod
		sums[i] = (sums[i-1] + dp[i]) % mod
	}

	return dp[n]
}
