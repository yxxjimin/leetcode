// Runtime : 22 ms
// Memory  : 14.28 MB
package feb25

func countBadPairs(nums []int) int64 {
	n := len(nums)
	diffCnts := make(map[int]int)
	good := 0

	for idx, num := range nums {
		diffCnts[num-idx]++
	}
	for _, diff := range diffCnts {
		good += diff * (diff - 1) / 2
	}

	return int64(n*(n-1)/2 - good)
}
