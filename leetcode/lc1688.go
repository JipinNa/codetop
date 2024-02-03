package leetcode

func numberOfMatches(n int) int {
	if n == 1 {
		return 0
	}
	next := n / 2
	if (n % 2) == 1 {
		return next + numberOfMatches(next+1)
	}
	return next + numberOfMatches(next)
}
