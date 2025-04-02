// Runtime : 0 ms
// Memory  : 4.22 MB
package apr25

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	maxVal := int64(0)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				maxVal = max(maxVal, (int64(nums[i]-nums[j]))*int64(nums[k]))
			}
		}
	}

	return maxVal
}
