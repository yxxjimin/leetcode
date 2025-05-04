// Runtime : 0 ms
// Memory  : 10.60 MB
package may25

func numEquivDominoPairs(dominoes [][]int) int {
	cnts := make(map[int]int)
	res := 0
	for _, pair := range dominoes {
		i, j := min(pair[0], pair[1]), max(pair[0], pair[1])
		res += cnts[i*10+j]
		cnts[i*10+j]++
	}

	return res
}
