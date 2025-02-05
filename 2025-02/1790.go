// Runtime : 0 ms
// Memory  : 3.91 MB
package feb25

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	front, back := 0, 0
	cnt := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			cnt++
			if cnt > 2 {
				return false
			} else if front == 0 {
				front = i
			} else {
				back = i
			}
		}
	}

	return s1[front] == s2[back] && s1[back] == s2[front]
}
