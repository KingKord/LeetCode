func numIslands(grid [][]byte) int {
	var (
		count = 0
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i,j,grid)
			}
		}
	}

	return count
}

func dfs(i,j int, grid [][]byte) {
	if i < 0 || i >= len(grid) || j >= len(grid[0]) || j < 0 {
		return
	}
	if grid[i][j] == '2' || grid[i][j] == '0' {
		return 
	}
	
	grid[i][j] = '2'

	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
	return
}