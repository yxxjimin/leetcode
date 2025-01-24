package main

// Runtime : 0 ms
// Memory  : 9.1 MB
// Comment :
//   - unsafe node if a path includes any cycle
//   - detect cycles with dfs
func isSafe(node int, graph [][]int, visited, path []bool) bool {
	if path[node] {
		return false
	} else if visited[node] {
		return true
	}

	visited[node] = true
	path[node] = true
	for _, adj := range graph[node] {
		if !isSafe(adj, graph, visited, path) {
			return false
		}
	}
	path[node] = false
	return true
}

func eventualSafeNodes(graph [][]int) []int {
	visited := make([]bool, len(graph))
	path := make([]bool, len(graph))

	for node := range graph {
		isSafe(node, graph, visited, path)
	}

	safe := make([]int, 0)
	for node, notSafe := range path {
		if !notSafe {
			safe = append(safe, node)
		}
	}
	return safe
}

// Runtime : 14 ms
// Memory  : 10.99 MB
// Comment :
//   - topological sorting; unvisited/unsorted nodes are unsafe nodes
func eventualSafeNodes2(graph [][]int) []int {
	inDegree := make([]int, len(graph))
	reversed := make([][]int, len(graph))
	safe := make([]bool, len(graph))

	for src, adj := range graph {
		for _, dst := range adj {
			reversed[dst] = append(reversed[dst], src)
			inDegree[src]++
		}
	}

	queue := []int{}
	for node, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		safe[node] = true

		for _, adj := range reversed[node] {
			inDegree[adj]--
			if inDegree[adj] == 0 {
				queue = append(queue, adj)
			}
		}
	}

	safeNodes := []int{}
	for node, isSafe := range safe {
		if isSafe {
			safeNodes = append(safeNodes, node)
		}
	}

	return safeNodes
}
