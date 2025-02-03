// Runtime : 0 ms
// Memory  : 4.38 MB
package feb25

func longestMonotonicSubarray(nums []int) int {
	state, maxLen, currLen := 0, 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			maxLen = max(maxLen, currLen)
		} else {
			if (nums[i]-nums[i-1])*state <= 0 {
				maxLen = max(maxLen, currLen)
				currLen = 1
			}
			currLen++
		}
		state = nums[i] - nums[i-1]
	}
	return max(maxLen, currLen)
}
