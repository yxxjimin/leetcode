// Runtime : 0 ms
// Memory  : 42.03 MB
package feb25

func tupleSameProduct(nums []int) int {
	products := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			products[nums[i]*nums[j]]++
		}
	}

	ans := 0
	for _, cnt := range products {
		ans += cnt * (cnt - 1) * 2 * 2
	}

	return ans
}
