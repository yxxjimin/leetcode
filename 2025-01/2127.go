// Runtime : 22 ms
// Memory  : 15.42 MB
// Comment :
//   - 위상 정렬로 사이클 찾아내기
//   - 길이 2짜리 사이클에 붙어있는 가장 긴 체인 연결해주기
//   - 길이가 2보다 큰 사이클들은 무조건 단일 사이클로만 구성
package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	degree := make([]int, n)
	for _, fav := range favorite {
		degree[fav]++
	}

	queue := []int{}
	for emp, deg := range degree {
		if deg == 0 {
			queue = append(queue, emp)
		}
	}

	depth := make([]int, n)
	for emp := range depth {
		depth[emp] = 1
	}

	for len(queue) > 0 {
		curr := queue[0]
		next := favorite[curr]
		queue = queue[1:]

		if d := depth[curr] + 1; d > depth[next] {
			depth[next] = d
		}

		degree[next]--
		if degree[next] == 0 {
			queue = append(queue, next)
		}
	}

	maxSingleCycle := 0
	maxMultiCycle := 0

	for emp, deg := range degree {
		if deg == 0 {
			continue
		}

		currCycle := 0
		curr := emp
		for degree[curr] != 0 {
			degree[curr] = 0
			currCycle++
			curr = favorite[curr]
		}

		if currCycle == 2 {
			maxMultiCycle += depth[emp] + depth[favorite[emp]]
		} else if currCycle > maxSingleCycle {
			maxSingleCycle = currCycle
		}
	}

	if maxSingleCycle > maxMultiCycle {
		return maxSingleCycle
	}
	return maxMultiCycle
}
