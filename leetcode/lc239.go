package leetcode

import "math"

func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}

func maxSlidingWindow(nums []int, k int) []int {
	count := len(nums) - k + 1
	ans := make([]int, count)
	var orderMax = func(start, end int) (int, int) {
		max := math.MinInt
		maxIndex := start
		for i := start; i <= end; i++ {
			if nums[i] >= max {
				max = nums[i]
				maxIndex = i
			}
		}
		return max, maxIndex
	}
	lastIndex := -1
	for i := 0; i < count; i++ {
		end := i + k - 1
		if i <= lastIndex && lastIndex <= end {
			if nums[end] > nums[lastIndex] {
				ans[i] = nums[end]
				lastIndex = end
			} else {
				ans[i] = nums[lastIndex]
			}
		} else {
			ans[i], lastIndex = orderMax(i, end)
		}
	}
	return ans
}
