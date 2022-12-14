package leetcode

func FindLength(nums1 []int, nums2 []int) int {
	return findLength(nums1, nums2)
}

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}
	max := 0
	for i := range nums1 {
		for j := range nums2 {
			if i == 0 && j == 0 {
				continue
			}
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return max
}
