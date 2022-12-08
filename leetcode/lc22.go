package leetcode

func generateParenthesis(n int) []string {
	var backtrack func(list []string, cur string, open, close int) []string
	backtrack = func(list []string, cur string, open, close int) []string {
		if len(cur) == 2*n {
			list = append(list, cur)
			return list
		}
		if open < n {
			list = backtrack(list, cur+"(", open+1, close)
		}
		if close < open {
			list = backtrack(list, cur+")", open, close+1)
		}
		return list
	}
	return backtrack([]string{}, "", 0, 0)
}
