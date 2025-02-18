// Runtime : 0 ms
// Memory  : 4.04 MB
package feb25

import "fmt"

func smallestNumber(pattern string) string {
	n := len(pattern)
	graph, deg := make(map[int][]int), make([]int, n+1)

	for i, r := range pattern {
		if r == 'I' {
			graph[i] = append(graph[i], i+1)
			deg[i+1]++
		} else {
			graph[i+1] = append(graph[i+1], i)
			deg[i]++
		}
	}

	queue := []int{}
	for i, d := range deg {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	idxOrder := []int{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		idxOrder = append(idxOrder, front+1)

		temp := []int{}
		for _, next := range graph[front] {
			deg[next]--
			if deg[next] == 0 {
				temp = append(temp, next)
			}
		}

		queue = append(temp, queue...)
	}

	var ans string
	for _, ord := range idxOrder {
		ans += fmt.Sprint(ord)
	}

	return ans
}
