package leetcode

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	l := 200
	for i := range strs {
		if len(strs[i]) < l {
			l = len(strs[i])
		}
	}
	for i := 1; i <= l; i++ {
		for j := range strs {
			if !strings.HasPrefix(strs[j], strs[0][:i]) {
				return commonPrefix
			}
		}
		commonPrefix = strs[0][:i]
	}
	return commonPrefix
}
