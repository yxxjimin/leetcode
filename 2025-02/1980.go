// Runtime : 0 ms
// Memory  : 4.17 MB
package feb25

import "slices"

func strToInt(s string) int {
	num := 0
	for i, r := range s {
		if r == '1' {
			num += 1 << (len(s) - 1 - i)
		}
	}
	return num
}

func intToStr(k, n int) string {
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		if r := k % 2; r == 1 {
			s[n-1-i] = '1'
		} else {
			s[n-1-i] = '0'
		}
		k /= 2
	}
	return string(s)
}

func findDifferentBinaryString(nums []string) string {
	appeared := []int{}
	for _, str := range nums {
		appeared = append(appeared, strToInt(str))
	}
	slices.Sort(appeared)

	min, n := 0, len(nums[0])
	for _, num := range appeared {
		if min < num {
			break
		} else {
			min++
		}
	}
	return intToStr(min, n)
}
