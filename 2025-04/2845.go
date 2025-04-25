// Runtime : 14 ms
// Memory  : 15.20 Mb
package apr25

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	ans := int64(0)
	prefix := 0
	cnt := make(map[int]int)
	cnt[0] = 1

	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefix++
		}
		ans += int64(cnt[(prefix-k+modulo)%modulo])
		cnt[prefix%modulo]++
	}

	return ans
}
