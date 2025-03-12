// Runtime : 0 ms
// Memory  : 7.9 MB
package mar25

func maximumCount(nums []int) int {
	n := len(nums)

	lo, hi := 0, n
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	lower := lo

	lo, hi = 0, n
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	upper := lo

	return max(lower, n-upper)
}
