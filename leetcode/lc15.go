package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		if nums[i] > 0 {
			break
		}
		a := nums[i]
		for l, r := i+1, len(nums)-1; l < r; {
			b, c := nums[l], nums[r]
			sum := a + b + c
			if sum == 0 {
				ret = append(ret, []int{a, b, c})
				for l < r && nums[l] == b {
					l++
				}
			} else if sum > 0 {
				for l < r && nums[r] == c {
					r--
				}
			} else {
				for l < r && nums[l] == b {
					l++
				}
			}
		}
		for i < len(nums) && nums[i] == a {
			i++
		}
	}
	return ret
}
