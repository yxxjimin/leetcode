// Runtime : 56 ms
// Memory  : 11.3 MB
// Comment:
//   - 숫자들 정렬 후 `limit`에 따라 그룹화
//   - 그룹별로 정렬 후 그룹 내 원소들끼리 스왑
package main

import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ans := make([]int, n)
	idx := make([]int, n)

	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int {
		return nums[i] - nums[j]
	})

	for grpStart := 0; grpStart < n; {
		grpEnd := grpStart + 1
		for grpEnd < n && nums[idx[grpEnd]]-nums[idx[grpEnd-1]] <= limit {
			grpEnd++
		}
		idxGrp := append([]int{}, idx[grpStart:grpEnd]...)
		slices.Sort(idxGrp)
		for k := range idxGrp {
			ans[idxGrp[k]] = nums[idx[grpStart+k]]
		}
		grpStart = grpEnd
	}

	return ans
}
