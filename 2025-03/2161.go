// Runtime : 17 ms
// Memory  : 21.57 MB
package mar25

func pivotArray(nums []int, pivot int) []int {
	less, equal, greater := []int{}, []int{}, []int{}

	for _, num := range nums {
		if num > pivot {
			greater = append(greater, num)
		} else if num < pivot {
			less = append(less, num)
		} else {
			equal = append(equal, num)
		}
	}

	ans := append(less, equal...)
	return append(ans, greater...)
}
