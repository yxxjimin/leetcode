// Runtime : 0 ms
// Memory  : 4.10 MB
package feb25

func getHappyString(n int, k int) string {
	getLo := map[rune]rune{
		'a': 'b', 'b': 'a', 'c': 'a',
	}
	getHi := map[rune]rune{
		'a': 'c', 'b': 'c', 'c': 'b',
	}

	lo, hi, ans := 1, 3*(1<<(n-1)), []rune{}
	if k > hi {
		return ""
	} else if mid2 := (1 << n); k > mid2 {
		lo = mid2 + 1
		ans = append(ans, 'c')
	} else if mid1 := (1 << (n - 1)); k > mid1 {
		lo, hi = mid1+1, mid2
		ans = append(ans, 'b')
	} else {
		hi = mid1
		ans = append(ans, 'a')
	}

	for lo < hi {
		mid := (lo + hi) / 2
		if k > mid {
			lo = mid + 1
			ans = append(ans, getHi[ans[len(ans)-1]])
		} else {
			hi = mid
			ans = append(ans, getLo[ans[len(ans)-1]])
		}
	}

	return string(ans)
}
