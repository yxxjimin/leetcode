// Runtime : 111 ms
// Memory  : 35.88 MB
package feb25

import "math"

func findBobPath(src, time int, path map[int]int, graph map[int][]int, visited []bool) bool {
	visited[src] = true
	path[src] = time

	if src == 0 {
		return true
	}

	for _, next := range graph[src] {
		if !visited[next] {
			if findBobPath(next, time+1, path, graph, visited) {
				return true
			}
		}
	}

	delete(path, src)
	return false
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n, maxIncome := len(amount), math.MinInt32
	graph := make(map[int][]int)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited, bobPath := make([]bool, n), make(map[int]int)
	findBobPath(bob, 0, bobPath, graph, visited)
	visited = make([]bool, n)

	queue, time := [][3]int{{0, 0}}, 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			front := queue[0]
			queue = queue[1:]
			node, income := front[0], front[1]
			visited[node] = true

			if bobTime, ok := bobPath[node]; !ok || time < bobTime {
				income += amount[node]
			} else if time == bobTime {
				income += amount[node] / 2
			}

			if len(graph[node]) == 1 && node != 0 {
				maxIncome = max(maxIncome, income)
			}

			for _, next := range graph[node] {
				if !visited[next] {
					queue = append(queue, [3]int{next, income})
				}
			}
		}
		time++
	}

	return maxIncome
}
