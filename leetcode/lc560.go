package leetcode

func SubarraySum(nums []int, k int) int {
	return subarraySum(nums, k)
}

func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	score := make(map[int]int)
	score[0] = 1
	for i := range nums {
		pre += nums[i]
		if v, ok := score[pre-k]; ok {
			count += v
		}
		score[pre]++
	}
	return count
}

func subarraySum1(nums []int, k int) int {
	pre, count := 0, 0
	score := make(map[int]int)
	score[0] = 1
	for i := range nums {
		pre += nums[i]
		if _, ok := score[pre-k]; ok {
			count++
		}
		score[pre] = 1
	}
	return count
}
