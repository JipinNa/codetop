package leetcode

import "math"

func maxProfit(prices []int) int {
	min, max := math.MaxInt, 0
	for _, p := range prices {
		if p < min {
			min = p
		}
		res := p - min
		if res > max {
			max = res
		}
	}
	return max
}
