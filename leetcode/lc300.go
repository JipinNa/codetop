package leetcode

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxL := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				sum := dp[j] + 1
				if sum > dp[i] {
					dp[i] = sum
				}
			}
		}
		if dp[i] > maxL {
			maxL = dp[i]
		}
	}
	return maxL
}
