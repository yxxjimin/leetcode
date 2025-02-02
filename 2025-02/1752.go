// Runtime : 0 ms
// Memory  : 4.08 MB
package feb25

func check(nums []int) bool {
	rotated, n := false, len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] && rotated {
			return false
		} else if nums[i] < nums[i-1] {
			rotated = true
		}
	}
	if rotated {
		return nums[n-1] <= nums[0]
	}
	return true
}
