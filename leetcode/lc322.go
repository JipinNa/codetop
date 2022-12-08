package leetcode

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < max; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
