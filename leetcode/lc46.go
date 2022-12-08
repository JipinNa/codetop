package leetcode

func Permute(nums []int) [][]int {
	return permute(nums)
}

func permute(nums []int) [][]int {
	length := len(nums)
	if length == 1 {
		return [][]int{nums}
	}
	ans := make([][]int, 0, factorial(length))
	for i := range nums {
		input := make([]int, 0, length-1)
		input = append(input, nums[:i]...)
		input = append(input, nums[i+1:]...)
		next := permute(input)
		for n := range next {
			ans = append(ans, append(next[n], nums[i]))
		}
	}
	return ans
}

func factorial(count int) int {
	factorial := 1
	for i := 1; i <= count; i++ {
		factorial *= i
	}
	return factorial
}
