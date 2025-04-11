// Runtime : 6 ms
// Memory  : 4.62 MB
package apr25

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for num := low; num < high+1; num++ {
		if 10 <= num && num < 100 && num%11 == 0 {
			ans++
		} else if 1000 <= num && num < 10000 {
			leftSum := num/1000 + (num/100)%10
			rightSum := (num/10)%10 + num%10
			if leftSum == rightSum {
				ans++
			}
		}
	}
	return ans
}
