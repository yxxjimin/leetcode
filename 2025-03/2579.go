// Runtime : 0 ms
// Memory  : 3.89 MB
package mar25

func coloredCells(n int) int64 {
	n64 := int64(n)
	return int64(1) + int64(2)*n64*(n64-1)
}
