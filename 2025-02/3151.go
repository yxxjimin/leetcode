// Runtime : 0 ms
// Memory  : 4.74 MB
package feb25

func isArraySpecial(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if (nums[i]^nums[i-1])&1 != 1 {
			return false
		}
	}
	return true
}
