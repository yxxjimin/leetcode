// Runtime : 201 ms
// Memory  : 30.39 MB
package mar25

import "slices"

func checkValidCuts(n int, rectangles [][]int) bool {
	horEnd, verEnd := -1, -1
	horGrps, verGrps := 0, 0

	hor, ver := make([][]int, 0), make([][]int, 0)

	for _, coord := range rectangles {
		x1, y1, x2, y2 := coord[0], coord[1], coord[2], coord[3]
		hor = append(hor, []int{x1, x2})
		ver = append(ver, []int{y1, y2})
	}

	slices.SortFunc(hor, func(a, b []int) int {
		return a[0] - b[0]
	})
	slices.SortFunc(ver, func(a, b []int) int {
		return a[0] - b[0]
	})

	for _, coord := range hor {
		start, end := coord[0], coord[1]
		if start >= horEnd {
			horGrps++
			horEnd = end
		} else {
			horEnd = max(horEnd, end)
		}
	}

	for _, coord := range ver {
		start, end := coord[0], coord[1]
		if start >= verEnd {
			verGrps++
			verEnd = end
		} else {
			verEnd = max(verEnd, end)
		}
	}

	return horGrps > 2 || verGrps > 2
}
