package leetcode

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

func subsets(nums []int) [][]int {
	ret := make([][]int, 0, factorial1(len(nums))+1)
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ret = append(ret, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ret
}

func factorial1(count int) int {
	sum := 1
	for i := 1; i <= count; i++ {
		e1 := 1
		e2 := 1
		for j := 0; j < i; j++ {
			e1 *= count - j
			e2 *= j + 1
		}
		sum += e1 / e2
	}
	return sum
}
