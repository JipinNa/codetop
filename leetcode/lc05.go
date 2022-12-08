package leetcode

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	maxLength := 1
	begin := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for length := 2; length <= l; length++ {
		for i := 0; i < l; i++ {
			j := length + i - 1
			if j >= l {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if length < 4 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && length > maxLength {
				maxLength = length
				begin = i
			}
		}
	}
	return s[begin : begin+maxLength]
}
