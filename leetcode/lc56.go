package leetcode

import "sort"

func Merge56(intervals [][]int) [][]int {
	return merge56(intervals)
}

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := [][]int{intervals[0]}
	for i := range intervals {
		prev := ans[len(ans)-1]
		if intervals[i][0] <= prev[1] {
			prev[1] = max(intervals[i][1], prev[1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
