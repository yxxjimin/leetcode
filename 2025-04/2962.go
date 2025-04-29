// Runtime : 1 ms
// Memory  : 13.07 MB
package apr25

func countSubarrays3(nums []int, k int) int64 {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}

	res := int64(0)
	lo, cnt := 0, 0

	for hi := range nums {
		if nums[hi] == maxVal {
			cnt++
		}
		for cnt == k {
			res += int64(len(nums) - hi)
			if nums[lo] == maxVal {
				cnt--
			}
			lo++
		}
	}

	return res
}
