package leetcode

func majorityElement(nums []int) int {
	num := nums[0]
	count := 0
	for i := range nums {
		if num == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				num = nums[i+1]
			}
		}
	}
	return num
}
