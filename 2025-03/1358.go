// Runtime : 3 ms
// Memory  : 5.75 MB
package mar25

func numberOfSubstrings(s string) int {
	ans := 0
	ap, bp, cp := -1, -1, -1

	for upper, ch := range s {
		if ch == 'a' {
			ap = upper
		} else if ch == 'b' {
			bp = upper
		} else {
			cp = upper
		}
		if ap > -1 && bp > -1 && cp > -1 {
			ans += min(ap, bp, cp) + 1
		}
	}

	return ans
}
