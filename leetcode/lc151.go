package leetcode

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	ret := ""
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == "" {
			continue
		}
		ret += " " + list[i]
	}
	return ret[1:]
}
