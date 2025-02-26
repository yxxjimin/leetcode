// Runtime : 0 ms
// Memory  : 4.06 MB
package feb25

func backtrack2(idx int, seq []int, used []bool) bool {
	n := len(used) - 1

	if idx == len(seq) {
		return true
	}
	if seq[idx] != 0 {
		return backtrack2(idx+1, seq, used)
	}

	for i := n; i > 0; i-- {
		if !used[i] {
			used[i] = true
			seq[idx] = i
			if i == 1 {
				if backtrack2(idx+1, seq, used) {
					return true
				}
			} else if idx+i < len(seq) && seq[idx+i] == 0 {
				seq[idx+i] = i
				if backtrack2(idx+1, seq, used) {
					return true
				}
				seq[idx+i] = 0
			}
			used[i] = false
			seq[idx] = 0
		}
	}
	return false
}

func constructDistancedSequence(n int) []int {
	seq, used := make([]int, 2*n-1), make([]bool, n+1)
	backtrack2(0, seq, used)
	return seq
}
