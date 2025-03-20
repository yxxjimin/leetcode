// Runtime : 1 ms
// Memory  : 7.19 MB
package mar25

func divideArray(nums []int) bool {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	for _, v := range cnt {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
