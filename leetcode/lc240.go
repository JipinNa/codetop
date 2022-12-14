package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := m-1, 0
	for x >= 0 && y < n {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			x--
			continue
		}
		if matrix[x][y] < target {
			y++
			continue
		}
	}
	return false
}
