package leetcode

func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
		mem[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}
		if mem[i][j] != 0 {
			return mem[i][j]
		}
		res := max(sum[j+1]-sum[i+1]-dfs(i+1, j), sum[j]-sum[i]-dfs(i, j-1))
		mem[i][j] = res
		return res
	}
	return dfs(0, n-1)
}

func stoneGameVII2(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
		dp[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(sum[j+1]-sum[i+1]-dp[i+1][j], sum[j]-sum[i]-dp[i][j-1])
		}
	}
	return dp[0][n-1]
}
