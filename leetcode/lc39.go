package leetcode

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	var backtrace func([]int, int, int)
	backtrace = func(combination []int, cur int, index int) {
		if cur > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			combination = append(combination, candidates[i])

			if cur+candidates[i] == target {
				ret = append(ret, append([]int(nil), combination...))
			} else {
				backtrace(combination, cur+candidates[i], i)
			}
			combination = combination[:len(combination)-1]
		}
	}
	backtrace([]int(nil), 0, 0)
	return ret
}
