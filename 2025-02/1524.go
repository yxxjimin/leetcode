// Runtime : 4 ms
// Memory  : 11.61 MB
package feb25

import "math"

func numOfSubarrays(arr []int) int {
	mod := int(math.Pow(10, 9)) + 7
	prefixSum := 0
	oddPrefixSumCnt, evenPrefixSumCnt := 0, 1
	cnt := 0

	for _, num := range arr {
		prefixSum += num
		if prefixSum&1 == 1 {
			oddPrefixSumCnt++
			cnt += evenPrefixSumCnt
		} else {
			evenPrefixSumCnt++
			cnt += oddPrefixSumCnt
		}
		cnt %= mod
	}

	return cnt
}
