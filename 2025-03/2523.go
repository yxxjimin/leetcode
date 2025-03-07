// Runtime : 65 ms
// Memory  : 10.41 MB
package mar25

import "math"

func fillComposites(upper int) []bool {
	isComposite := make([]bool, upper+1)

	for i := 2; i*i < upper+1; i++ {
		if !isComposite[i] {
			for k := i * i; k < upper+1; k += i {
				isComposite[k] = true
			}
		}
	}

	return isComposite
}

func closestPrimes(left int, right int) []int {
	isComposite := fillComposites(right)
	prev, curr := -1, -1
	diff, ans := math.MaxInt32, []int{prev, curr}

	for i := left; i < right+1; i++ {
		if !isComposite[i] {
			prev, curr = curr, i
			if newDiff := curr - prev; newDiff < diff && prev > 1 {
				diff = newDiff
				ans[0], ans[1] = prev, curr
			}
		}
	}

	return ans
}
