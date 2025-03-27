// Runtime : 32 ms
// Memory  : 11.54 MB
package mar25

func minimumIndex(nums []int) int {
	n := len(nums)
	cnts := make(map[int]int)
	mode := nums[0]

	for _, num := range nums {
		cnts[num]++
		if cnts[num] > cnts[mode] {
			mode = num
		}
	}

	modeCnt := make([]int, n)
	curr := 0

	for i, num := range nums {
		if num == mode {
			curr++
		}
		modeCnt[i] = curr
	}

	for i := 0; i < n-1; i++ {
		if modeCnt[i] > (i+1)/2 && modeCnt[n-1]-modeCnt[i] > (n-i-1)/2 {
			return i
		}
	}

	return -1
}
