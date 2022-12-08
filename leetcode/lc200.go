package leetcode

func NumIslands(grid [][]byte) int {
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
	r := len(grid)
	c := len(grid[0])
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= r {
			return
		}
		if j < 0 || j >= c {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	ret := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				ret++
				dfs(grid, i, j)
			}
		}
	}
	return ret
}
