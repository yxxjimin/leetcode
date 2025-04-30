// Runtime : 0 ms
// Memory  : 4.70 MB
package apr25

func isEvenDigit(num int) bool {
	cnt := 1
	for num > 0 {
		if num/10 > 0 {
			cnt++
		}
		num /= 10
	}
	return cnt%2 == 0
}

func findNumbers(nums []int) int {
	cnt := 0
	for _, num := range nums {
		if isEvenDigit(num) {
			cnt++
		}
	}
	return cnt
}
