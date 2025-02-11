// Runtime : 0 ms
// Memory  : 4.10 MB
package feb25

func removeOccurrences(s string, part string) string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= len(part) && string(stack[len(stack)-len(part):]) == part {
			stack = stack[:len(stack)-len(part)]
		}
	}

	return string(stack)
}
