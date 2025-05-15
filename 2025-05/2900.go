// Runtime : 0 ms
// Memory  : 5.63 MB
package may25

func getLongestSubsequence(words []string, groups []int) []string {
	ans := words[:1]
	n := len(words)

	for i := range n - 1 {
		if groups[i] != groups[i+1] {
			ans = append(ans, words[i+1])
		}
	}

	return ans
}
