package leetcode

func trap(height []int) int {
	l := len(height)
	leftDp := make([]int, l)
	rightDp := make([]int, l)
	leftDp[0] = height[0]
	rightDp[l-1] = height[l-1]

	for i, j := 1, l-2; i < l && j >= 0; {
		if height[i] > leftDp[i-1] {
			leftDp[i] = height[i]
		} else {
			leftDp[i] = leftDp[i-1]
		}
		if height[j] > rightDp[j+1] {
			rightDp[j] = height[j]
		} else {
			rightDp[j] = rightDp[j+1]
		}
		i++
		j--
	}
	var ans int
	for i := range height {
		if leftDp[i] > rightDp[i] {
			ans += rightDp[i] - height[i]
		} else {
			ans += leftDp[i] - height[i]
		}
	}
	return ans
}
