// Runtime : 0 ms
// Memory  : 4.65 MB
package apr25

func numRabbits(answers []int) int {
	cnts := make(map[int]int)
	total := 0

	for _, ans := range answers {
		if _, ok := cnts[ans]; ok {
			if cnts[ans] == ans {
				total += ans + 1
				cnts[ans] = 0
			} else {
				cnts[ans]++
			}
		} else {
			cnts[ans] = 0
		}
	}

	for k, _ := range cnts {
		total += k + 1
	}

	return total
}
