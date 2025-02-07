// Runtime : 46 ms
// Memory  : 31.38 MB
package feb25

func queryResults(_ int, queries [][]int) []int {
	colorOf := make(map[int]int)
	labelsCntOf := make(map[int]int)
	result := make([]int, len(queries))

	for i, query := range queries {
		label, color := query[0], query[1]

		if prev, exists := colorOf[label]; exists {
			labelsCntOf[prev]--
			if labelsCntOf[prev] == 0 {
				delete(labelsCntOf, prev)
			}
		}

		colorOf[label] = color
		labelsCntOf[color]++

		result[i] = len(labelsCntOf)
	}

	return result
}
