// Runtime : 0 ms
// Memory  : 4.10 MB
package feb25

func clearDigits(s string) string {
	ans := ""
	for _, c := range s {
		if 48 <= c && c < 58 {
			ans = ans[:len(ans)-1]
		} else {
			ans += string(c)
		}
	}

	return ans
}
