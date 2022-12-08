package leetcode

import (
	"bytes"
)

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	l := len(s)
	for i := 0; i < l; i++ {
		state := bytes.Repeat([]byte{0}, 12)
		innerMax := 0
		for j := i; j < l; j++ {
			pos := (s[j] - 32) / 8
			bit := (s[j] - 32) % 8
			a := byte(1 << bit)
			if (state[pos] & a) == a {
				break
			}
			state[pos] |= 1 << bit
			innerMax++
		}
		if innerMax > max {
			max = innerMax
		}
	}
	return max
}
