// Runtime : 0 ms
// Memory  : 12.88 MB
package mar25

func longestNiceSubarray(nums []int) int {
	sum := 0
	lower := 0
	maxLength := 0

	for i, num := range nums {
		if sum|num != sum+num {
			maxLength = max(maxLength, i-lower)
			for sum|num != sum+num {
				sum -= nums[lower]
				lower++
			}
		}
		sum += num
	}

	return max(maxLength, len(nums)-lower)
}
