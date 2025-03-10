// Runtime : 126 ms
// Memory  : 9.40 MB
package mar25

func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func allVowelsChecked(vowels map[rune]int) bool {
	for _, v := range vowels {
		if v <= 0 {
			return false
		}
	}
	return true
}

func countOfSubstrings(word string, k int) int64 {
	ans := int64(0)
	lower, upper := 0, 0
	l := len(word)

	vowelCnt := make(map[rune]int)
	consonantCnt := 0

	nextConsonant := make([]int, l)
	nextConsonantIdx := l
	for i := l - 1; i >= 0; i-- {
		nextConsonant[i] = nextConsonantIdx
		if !isVowel(rune(word[i])) {
			nextConsonantIdx = i
		}
	}

	for upper < l {
		if ch := rune(word[upper]); isVowel(ch) {
			vowelCnt[ch]++
		} else {
			consonantCnt++
		}

		for consonantCnt > k && lower < l {
			if ch := rune(word[lower]); isVowel(ch) {
				vowelCnt[ch]--
				if vowelCnt[ch] == 0 {
					delete(vowelCnt, ch)
				}
			} else {
				consonantCnt--
			}
			lower++
		}

		for lower < l && len(vowelCnt) == 5 && consonantCnt == k {
			ans += int64(nextConsonant[upper] - upper)
			if ch := rune(word[lower]); isVowel(ch) {
				vowelCnt[ch]--
				if vowelCnt[ch] == 0 {
					delete(vowelCnt, ch)
				}
			} else {
				consonantCnt--
			}
			lower++
		}
		upper++
	}
	return ans
}
