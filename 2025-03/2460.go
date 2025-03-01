// Runtime : 0 ms
// Memory  : 5.00 MB
package mar25

func applyOperations(nums []int) []int {
	n := len(nums)
	ans := []int{}

	for i := 0; i < n-1; i++ {
		if nums[i] != 0 {
			if nums[i] == nums[i+1] {
				nums[i] *= 2
				nums[i+1] = 0
			}
			ans = append(ans, nums[i])
		}
	}
	ans = append(ans, nums[n-1])
	ans = append(ans, make([]int, n-len(ans))...)

	return ans
}
