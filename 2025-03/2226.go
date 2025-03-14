// Runtime : 6 ms
// Memory  : 10.16 MB
package mar25

func canAllocate(candies []int, k, val int64) bool {
	children := int64(0)
	for _, pile := range candies {
		children += int64(pile) / val
	}
	return children >= k
}

func maximumCandies(candies []int, k int64) int {
	totalCnt := int64(0)
	for _, cnt := range candies {
		totalCnt += int64(cnt)
	}

	if totalCnt < k {
		return 0
	}

	lo, hi := int64(1), totalCnt/k
	for lo < hi {
		mid := (lo + hi) / 2
		if canAllocate(candies, k, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if canAllocate(candies, k, lo) {
		return int(lo)
	}
	return int(lo - 1)
}
