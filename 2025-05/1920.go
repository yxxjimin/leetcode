// Runtime : 0 ms
// Memory  : 8.58 MB
package may25

func buildArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := range n {
		ans[i] = nums[nums[i]]
	}

	return ans
}
