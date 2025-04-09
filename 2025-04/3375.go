// Runtime : 7 ms
// Memory  : 5.86 MB
package apr25

func minOperations(nums []int, k int) int {
	keys := make(map[int]bool)
	minKey := 100

	for _, num := range nums {
		keys[num] = true
		minKey = min(minKey, num)
	}

	if minKey < k {
		return -1
	}

	keys[k] = true

	return len(keys) - 1
}
