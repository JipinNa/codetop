package leetcode

import "math"

func FindPeakElement(nums []int) int {
	return findPeakElement(nums)
}

func findPeakElement(nums []int) int {
	var find = func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}
	l, r := 0, len(nums)-1
	for {
		mid := (l + r) / 2
		if find(mid) > find(mid-1) && find(mid) > find(mid+1) {
			return mid
		}
		if find(mid) < find(mid+1) {
			l = mid
		} else {
			r = mid
		}
	}
}
