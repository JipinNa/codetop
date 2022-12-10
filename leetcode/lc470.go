package leetcode

func rand10() int {
	a, b := rand7(), rand7()
	for a == 1 {
		a = rand7()
	}
	for b > 5 {
		b = rand7()
	}
	if a&1 == 1 {
		return b
	}
	return 5 + b
}

func rand7() int {
	return 0
}
