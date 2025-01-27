package main

// Runtime : 11 ms
// Memory  : 9.27 MB
func merge(from, to []bool) {
	for i := range from {
		to[i] = from[i] || to[i]
	}
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	required := make([][]bool, numCourses)
	for i := range required {
		required[i] = make([]bool, numCourses)
	}

	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, req := range prerequisites {
		prev, next := req[0], req[1]
		graph[prev] = append(graph[prev], next)
		indegree[next]++
		required[next][prev] = true
	}

	queue := []int{}
	for num, deg := range indegree {
		if deg == 0 {
			queue = append(queue, num)
		}
	}

	for len(queue) > 0 {
		prev := queue[0]
		queue = queue[1:]
		for _, next := range graph[prev] {
			indegree[next]--
			merge(required[prev], required[next])
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	answer := make([]bool, len(queries))
	for i, query := range queries {
		answer[i] = required[query[1]][query[0]]
	}

	return answer
}

// Runtime : 4 ms
// Memory  : 9.59 MB
// Comment :
//   - 불리언 슬라이스를 비트 연산자로 구현한 코드가 있길래 따라해봄
type HashSet struct {
	low  uint64 // Encode membership from [0, 63]
	high uint64 // Encode membership from [64, 100]
}

func (hs *HashSet) add(e int) {
	if e < 64 {
		hs.low |= 1 << e
	} else {
		hs.high |= 1 << (e - 64)
	}
}

func (hs *HashSet) contains(e int) bool {
	if e < 64 {
		return (hs.low & (1 << e)) != 0
	} else {
		return (hs.high & (1 << (e - 64))) != 0
	}
}

func (hs *HashSet) union(other *HashSet) {
	hs.low |= other.low
	hs.high |= other.high
}

func checkIfPrerequisite2(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	required := make([]*HashSet, numCourses)
	for i := range required {
		required[i] = &HashSet{0, 0}
	}

	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, req := range prerequisites {
		prev, next := req[0], req[1]
		graph[prev] = append(graph[prev], next)
		indegree[next]++
		required[next].add(prev)
	}

	queue := []int{}
	for num, deg := range indegree {
		if deg == 0 {
			queue = append(queue, num)
		}
	}

	for len(queue) > 0 {
		prev := queue[0]
		queue = queue[1:]
		for _, next := range graph[prev] {
			indegree[next]--
			required[next].union(required[prev])
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	answer := make([]bool, len(queries))
	for i, query := range queries {
		answer[i] = required[query[1]].contains(query[0])
	}

	return answer
}
