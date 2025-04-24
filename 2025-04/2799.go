// Runtime : 3 ms
// Memory  : 8.11 MB
package apr25

func countCompleteSubarrays(nums []int) int {
	n := len(nums)
	elems := make(map[int]bool)
	for _, num := range nums {
		elems[num] = true
	}
	distincts := len(elems)

	curr := make(map[int]int)
	lo, hi := 0, distincts-1
	total := 0

	for i := lo; i < hi+1; i++ {
		curr[nums[i]]++
	}

	for lo <= n-distincts && hi < n {
		if len(curr) == distincts {
			total += n - hi
			curr[nums[lo]]--
			if curr[nums[lo]] == 0 {
				delete(curr, nums[lo])
			}
			lo++
		} else {
			hi++
			if hi < n {
				curr[nums[hi]]++
			}
		}
	}

	return total
}
