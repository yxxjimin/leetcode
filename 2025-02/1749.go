// Runtime : 0 ms
// Memory  : 9.84 MB
package feb25

func maxAbsoluteSum(nums []int) int {
	sum, minSum, maxSum := 0, 0, 0

	for _, num := range nums {
		sum += num
		minSum, maxSum = min(minSum, sum), max(maxSum, sum)
	}

	return maxSum - minSum
}
