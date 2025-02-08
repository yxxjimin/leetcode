// Runtime : 118 ms
// Memory  : 50.12 MB
package feb25

import "sort"

type NumberContainers struct {
	numOf  map[int]int
	idxOf  map[int][]int
	sorted map[int]bool
}

func Constructor() NumberContainers {
	return NumberContainers{
		numOf:  make(map[int]int),
		idxOf:  make(map[int][]int),
		sorted: make(map[int]bool),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if prev, exists := this.numOf[index]; exists {
		if prev != number {
			for i := 0; i < len(this.idxOf[prev]); i++ {
				if this.idxOf[prev][i] == index {
					this.idxOf[prev] = append(this.idxOf[prev][:i], this.idxOf[prev][i+1:]...)
				}
			}
			this.idxOf[number] = append(this.idxOf[number], index)
			this.sorted[number] = false
		}
	} else {
		this.idxOf[number] = append(this.idxOf[number], index)
		this.sorted[number] = false
	}

	this.numOf[index] = number
}

func (this *NumberContainers) Find(number int) int {
	if len(this.idxOf[number]) == 0 {
		return -1
	}
	if !this.sorted[number] {
		sort.Ints(this.idxOf[number])
		this.sorted[number] = true
	}
	return this.idxOf[number][0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
