package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var find func(i, j int) int
	find = func(i, j int) int {
		if i < 0 || i >= len(grid) {
			return 0
		}
		if j < 0 || j >= len(grid[0]) {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		ret := 1
		ret += find(i-1, j)
		ret += find(i+1, j)
		ret += find(i, j-1)
		ret += find(i, j+1)
		return ret
	}
	max := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				m := find(i, j)
				if m > max {
					max = m
				}
			}
		}
	}
	return max
}
