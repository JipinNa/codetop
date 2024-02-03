package leetcode

func countConsistentStrings(allowed string, words []string) int {
	res := 0
	flag := 0
	for _, c := range allowed {
		flag |= 1 << (c - 'a')
	}
	var flag1 int
	for i, _ := range words {
		flag1 = 0
		for _, c := range words[i] {
			flag1 |= 1 << (c - 'a')
		}
		if (flag1 | flag) == flag {
			res++
		}
	}
	return res
}
