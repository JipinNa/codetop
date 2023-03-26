package leetcode

func dailyTemperatures(temperatures []int) []int {
	// 单调栈
	// 什么时候使用单调栈？
	// 当需要找到左边或者右边第一个比当前位置大或者小就可以使用
	ret := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ret
}
