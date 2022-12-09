package leetcode

import "math"

func MinPathSum(grid [][]int) int {
	return minPathSum(grid)
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := range dp {
		for j := range dp[i] {
			if i == 0 && j == 0 {
				continue
			}
			minDp := math.MaxInt
			if i > 0 {
				minDp = min(minDp, dp[i-1][j])
			}
			if j > 0 {
				minDp = min(minDp, dp[i][j-1])
			}
			dp[i][j] = minDp + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
