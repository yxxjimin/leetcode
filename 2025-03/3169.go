// Runtime : 23 ms
// Memory  : 19.84 MB
package mar25

import "slices"

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})

	end, available := 0, 0

	for _, interval := range meetings {
		nextStart, nextEnd := interval[0], interval[1]
		if gap := nextStart - end; gap > 0 {
			available += gap - 1
		}
		end = max(end, nextEnd)
	}

	available += days - end

	return available
}
