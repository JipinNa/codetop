package leetcode

func MaximalSquare(matrix [][]byte) int {
	return maximalSquare(matrix)
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxL := 0
	for i := range matrix {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxL = 1
		}
	}
	for j := range matrix[0] {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxL = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				maxL = max(maxL, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maxL * maxL
}
