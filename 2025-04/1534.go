// Runtime : 3 ms
// Memory  : 4.17 MB
package apr25

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	n := len(arr)
	for i := 0; i < n-2; i++ {
		first := arr[i]
		for j := i + 1; j < n-1; j++ {
			second := arr[j]
			if abs(second-first) <= a {
				for k := j + 1; k < n; k++ {
					third := arr[k]
					if abs(third-second) <= b && abs(first-third) <= c {
						ans++
					}
				}
			}
		}
	}
	return ans
}
