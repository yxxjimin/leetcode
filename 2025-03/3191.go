// Runtime : 5 ms
// Memory  : 10.32 MB
package mar25

func minOperations(nums []int) int {
	n := len(nums)
	flips := 0

	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			flips++
		}
	}

	if nums[n-2]+nums[n-1] == 2 {
		return flips
	}
	return -1
}
