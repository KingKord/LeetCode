func orangesRotting(grid [][]int) int {
	var (
		secs    = -1
		nowFrom [][2]int
		counter int
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				nowFrom = append(nowFrom, [2]int{i, j})
			} else if grid[i][j] == 1 {
				counter++
			}
		}
	}
	
	if counter == 0 {
		return counter
	}
	for len(nowFrom) != 0 {
		var new [][2]int
		secs++
		for _, coordinates := range nowFrom {
			i, j := coordinates[0], coordinates[1]
			new = rotten(grid, i, j, new)
		}
		nowFrom = new
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	
	return secs
}

func rotten(grid [][]int, i, j int, coordinates [][2]int) [][2]int {
	if i-1 >= 0 && grid[i-1][j] == 1 {
		grid[i-1][j] = 2
		coordinates = append(coordinates, [2]int{i - 1, j})
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		grid[i+1][j] = 2
		coordinates = append(coordinates, [2]int{i + 1, j})
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		grid[i][j-1] = 2
		coordinates = append(coordinates, [2]int{i, j - 1})
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		grid[i][j+1] = 2
		coordinates = append(coordinates, [2]int{i, j + 1})
	}

	return coordinates
}
