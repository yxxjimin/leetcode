// Runtime : 0 ms
// Memory  : 10.04 MB
package apr25

func countSubarrays2(nums []int, k int64) int64 {
	n := len(nums)
	res, sum := int64(0), int64(0)
	for i, j := 0, 0; j < n; j++ {
		sum += int64(nums[j])
		for i <= j && sum*int64(j-i+1) >= k {
			sum -= int64(nums[i])
			i++
		}
		res += int64(j - i + 1)
	}
	return res
}
