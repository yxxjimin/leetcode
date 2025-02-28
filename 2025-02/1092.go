// Runtime : 629 ms
// Memory  : 34.37 MB
package feb25

func shortestCommonSupersequence(str1 string, str2 string) string {
	prevRow := make([]string, len(str2)+1)
	for i := range prevRow {
		prevRow[i] = str2[:i]
	}

	for i := 1; i < len(str1)+1; i++ {
		currRow := make([]string, len(str2)+1)
		currRow[0] = str1[:i]

		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				currRow[j] = prevRow[j-1] + string(str1[i-1])
			} else {
				if len(currRow[j-1]) < len(prevRow[j]) {
					currRow[j] = currRow[j-1] + string(str2[j-1])
				} else {
					currRow[j] = prevRow[j] + string(str1[i-1])
				}
			}
		}
		prevRow = currRow
	}

	return prevRow[len(str2)]
}
