func maxAreaOfIsland(grid [][]int) int {
	m := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				m = max(m, dfs(grid, i, j, 0))
			}
		}
	}
	
	return m
}

func dfs(grid [][]int, i, j, count int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 || grid[i][j] == 2 {
		return count
	}
	
	grid[i][j] = 2
	
	count++
	
	count = dfs(grid, i-1, j, count)
	count = dfs(grid, i+1, j, count)
	count = dfs(grid, i, j-1, count)
	count = dfs(grid, i, j+1, count)
	
	return count
}
