package leetcode

func climbStairs(n int) int {
	a, b := 1, 1
	for i := n - 2; i >= 0; i-- {
		a, b = b, a+b
	}
	return b
}

func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1]
		dp[i] += dp[i-2]
	}
	return dp[n]
}
