// Runtime : 0 ms
// Memory  : 5.61 MB
package mar25

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	ans, k := make([][]int, len(nums1)+len(nums2)), 0

	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i >= len(nums1) {
			ans[k] = nums2[j]
			k++
			j++
		} else if j >= len(nums2) {
			ans[k] = nums1[i]
			k++
			i++
		} else if nums1[i][0] < nums2[j][0] {
			ans[k] = nums1[i]
			k++
			i++
		} else if nums1[i][0] > nums2[j][0] {
			ans[k] = nums2[j]
			k++
			j++
		} else {
			ans[k] = []int{nums1[i][0], nums1[i][1] + nums2[j][1]}
			k++
			i++
			j++
		}
	}

	return ans[:k]
}
