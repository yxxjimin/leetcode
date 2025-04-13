// Runtime : 0 ms
// Memory  : 4.04 MB
package apr25

func pow(x, y int64) int64 {
	mod := int64(1000000007)
	ans := int64(1)
	for y > 0 {
		if y%2 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		y /= 2
	}
	return ans
}

func countGoodNumbers(n int64) int {
	mod := int64(1000000007)
	return int(pow(5, (n+1)/2) * pow(4, n/2) % mod)
}
