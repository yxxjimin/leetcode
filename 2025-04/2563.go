// Runtime : 30 ms
// Memory  : 10.09 MB
package apr25

import "sort"

func bisectLeft(a []int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func bisectRight(a []int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] > x {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	cnt := int64(0)

	for i, num := range nums {
		lo := max(lower-num, num)
		hi := upper - num
		if lo <= hi && i < len(nums)-1 {
			loi := bisectLeft(nums[i+1:], lo)
			hii := bisectRight(nums[i+1:], hi)
			cnt += int64(hii - loi)
		}
	}

	return cnt
}
