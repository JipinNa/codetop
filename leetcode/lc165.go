package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1l := strings.Split(version1, `.`)
	v2l := strings.Split(version2, `.`)
	i, j := 0, 0
	for i < len(v1l) || j < len(v2l) {
		v1, v2 := 0, 0
		if i < len(v1l) {
			v1, _ = strconv.Atoi(v1l[i])
		}
		if j < len(v2l) {
			v2, _ = strconv.Atoi(v2l[j])
		}
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
		i++
		j++
	}
	return 0
}
