package leetcode

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	stack := make([]int32, 0, 1)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if c == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if c == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if c == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
