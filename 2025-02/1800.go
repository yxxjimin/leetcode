// Runtime : 0 ms
// Memory  : 4.05 MB
package feb25

func maxAscendingSum(nums []int) int {
	maxSum, currSum := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currSum += nums[i]
		} else {
			maxSum = max(maxSum, currSum)
			currSum = nums[i]
		}
	}
	return max(maxSum, currSum)
}
