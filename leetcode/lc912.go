package leetcode

import "sort"

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low < high {
		l, r := low, high
		pivot := nums[l]
		for l < r {
			for l < r && nums[r] >= pivot {
				r--
			}
			nums[l] = nums[r]
			for l < r && nums[l] <= pivot {
				l++
			}
			nums[r] = nums[l]
		}
		nums[l] = pivot
		quickSort(nums, low, l-1)
		quickSort(nums, l+1, high)
	}
}

func partition(nums []int, low, high int) int {
	return 0
}
