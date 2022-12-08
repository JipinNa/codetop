package leetcode

import "math"

func MinWindow(s string, t string) string {
	return minWindow(s, t)
}

func minWindow(s string, t string) string {
	// sliding window
	hash := make(map[uint8]int, 0)
	for _, c := range t {
		hash[uint8(c)] = 0
	}
	var allExist = func() bool {
		for _, v := range hash {
			if v == 0 {
				return false
			}
		}
		return true
	}
	var count = func(key uint8, count int) {
		if _, ok := hash[key]; ok {
			hash[key] += count
		}
	}
	minLen, minR, minL := 0, 0, math.MaxInt
	count(s[0], 1)
	for l, r := 0, 0; r < len(s) && l <= r; {
		if allExist() {
			if minL-minR+1 < minLen {
				minL, minR = l, r
				minLen = minL - minR + 1
			}
			count(s[l], -1)
			l++
		} else {
			r++
			if r < len(s) {
				count(s[r], 1)
			}
		}
	}
	return s[minL:minR]
}
