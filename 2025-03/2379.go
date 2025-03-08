// Runtime : 0 ms
// Memory  : 3.94 MB
package mar25

func countWhiteBlocks(blocks string) int {
	cnt := 0
	for _, r := range blocks {
		if r == 'W' {
			cnt++
		}
	}
	return cnt
}

func minimumRecolors(blocks string, k int) int {
	ans := k
	for i := k; i < len(blocks)+1; i++ {
		ans = min(ans, countWhiteBlocks(blocks[i-k:i]))
	}

	return ans
}
