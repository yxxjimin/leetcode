// Runtime : 25 ms
// Memory  : 10.89 MB
package mar25

import "slices"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	adjSums := make([]int, n-1)
	minVal, maxVal := int64(0), int64(0)

	for i := range adjSums {
		adjSums[i] = weights[i] + weights[i+1]
	}

	slices.Sort(adjSums)

	for i := range k - 1 {
		minVal += int64(adjSums[i])
		maxVal += int64(adjSums[n-2-i])
	}

	return maxVal - minVal
}
