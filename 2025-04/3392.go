// Runtime : 0 ms
// Memory  : 6.55 MB
package apr25

func countSubarrays(nums []int) int {
	n := len(nums)
	res := 0
	for i := range n - 2 {
		if (nums[i]+nums[i+2])*2 == nums[i+1] {
			res++
		}
	}
	return res
}
