// Runtime : 0 ms
// Memory  : 3.87 MB
package mar25

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if r := n % 3; r == 2 {
			return false
		} else if r == 1 {
			n -= 1
		}
		n /= 3
	}
	return true
}
