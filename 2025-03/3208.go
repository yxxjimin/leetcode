// Runtime : 81 ms
// Memory  : 10.23 MB
package mar25

func numberOfAlternatingGroups(colors []int, k int) int {
	n, ans := len(colors), 0
	start := 0

	for i := 0; i < n+k-1; i++ {
		if colors[i%n] == colors[(i+1)%n] {
			grpLen := i - start + 1
			if grpLen >= k {
				ans += grpLen - k + 1
			}
			start = i + 1
		}
	}

	if grpLen := n + k - 1 - start; grpLen >= k {
		ans += grpLen - k + 1
	}

	return ans
}
