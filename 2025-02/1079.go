// Runtime : 10 ms
// Memory  : 3.88 MB
package feb25

func backtrack(src rune, cnt map[rune]int) int {
	if cnt[src] == 0 {
		return 0
	}

	ans := 1
	cnt[src]--
	for k := range cnt {
		ans += backtrack(k, cnt)
	}
	cnt[src]++
	return ans
}

func numTilePossibilities(tiles string) int {
	numCnt := make(map[rune]int)
	for _, ch := range tiles {
		numCnt[ch]++
	}

	ans := 0
	for k := range numCnt {
		ans += backtrack(k, numCnt)
	}

	return ans
}
