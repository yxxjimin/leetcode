// Runtime : 0 ms
// Memory  : 4.55 MB
package apr25

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	curr := []byte{}
	ch := prev[0]
	cnt := 1

	for i := 1; i < len(prev); i++ {
		if prev[i] == ch {
			cnt++
		} else {
			curr = append(curr, []byte(strconv.Itoa(cnt))...)
			curr = append(curr, ch)
			cnt = 1
			ch = prev[i]
		}
	}
	curr = append(curr, []byte(strconv.Itoa(cnt))...)
	curr = append(curr, ch)

	return string(curr)
}
