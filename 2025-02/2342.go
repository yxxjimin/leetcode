// Runtime : 30 ms
// Memory  : 10.10 MB
package feb25

func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func maximumSum(nums []int) int {
	digitSumToMaxNum := make(map[int]int)
	ans := -1

	for _, num := range nums {
		sum := digitSum(num)
		if maxNum := digitSumToMaxNum[sum]; maxNum > 0 {
			ans = max(ans, maxNum+num)
		}
		digitSumToMaxNum[sum] = max(digitSumToMaxNum[sum], num)
	}

	return ans
}
