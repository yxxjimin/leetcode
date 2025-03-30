// Runtime : 0 ms
// Memory  : 4.56 MB
package mar25

import "slices"

func partitionLabels(s string) []int {
	intervalMap := make(map[rune][]int)
	intervals := make([][]int, 0)

	for i, r := range s {
		if interval, ok := intervalMap[r]; !ok {
			intervalMap[r] = []int{i, i}
		} else {
			intervalMap[r][1] = max(interval[1], i)
		}
	}

	for _, interval := range intervalMap {
		intervals = append(intervals, interval)
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	prevStart, prevEnd := 0, 0
	answer := make([]int, 0)

	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if start > prevEnd {
			answer = append(answer, prevEnd-prevStart+1)
			prevStart, prevEnd = start, end
		} else {
			prevEnd = max(prevEnd, end)
		}
	}

	answer = append(answer, prevEnd-prevStart+1)

	return answer
}
