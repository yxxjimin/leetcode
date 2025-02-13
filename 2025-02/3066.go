// Runtime : 119 ms
// Memory  : 14.73 MB
package feb25

import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func minOperations(nums []int, k int) int {
	pq, cnt := PriorityQueue(nums), 0
	heap.Init(&pq)

	for len(pq) >= 2 && pq[0] < k {
		x, y := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		z := 2*x + y
		heap.Push(&pq, z)
		cnt++
	}

	return cnt
}
