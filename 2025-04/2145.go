// Runtime : 0 ms
// Memory  : 10.81 MB
package apr25

func numberOfArrays(differences []int, lower int, upper int) int {
	lb, ub := 0, 0
	curr := 0

	for _, diff := range differences {
		curr += diff
		lb = min(lb, curr)
		ub = max(ub, curr)

		if (ub - lb) > (upper - lower) {
			return 0
		}
	}

	return (upper - lower) - (ub - lb) + 1
}
