package leetcode

func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	var binary func(left, right int) int
	binary = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binary(left, mid-1)
		} else {
			return binary(mid+1, right)
		}
	}
	return binary(0, len(nums)-1)
}
