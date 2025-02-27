// Runtime : 15 ms
// Memory  : 4.56 MB
package feb25

import "strconv"

func substringSum(str []rune, tgt int) bool {
	if len(str) == 0 {
		return tgt == 0
	}

	offset := 0
	for i := range str {
		offset = offset*10 + int(str[i]-'0')
		if offset <= tgt && substringSum(str[i+1:], tgt-offset) {
			return true
		}
	}
	return false
}

func punishmentNumber(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if substringSum([]rune(strconv.Itoa(i*i)), i) {
			sum += i * i
		}
	}
	return sum
}
