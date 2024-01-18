func solve(board [][]byte) {
	visited := make(map[[2]int]struct{})
	lastJ := len(board[0]) - 1
	for i := range board {
		if board[i][0] == 'O' {
			dfs(board, i, 0, visited)
		}
		if board[i][lastJ] == 'O' {
			dfs(board, i, lastJ, visited)
		}
	}
	for j := range board[0] {
		if board[0][j] == 'O' {
			dfs(board, 0, j, visited)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(board, len(board)-1, j, visited)
		}
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			_, ok := visited[[2]int{i, j}]
			if !ok && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func dfs(board [][]byte, i, j int, visited map[[2]int]struct{}) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' {
		return
	}
	coordinates := [2]int{i, j}
	_, ok := visited[coordinates]
	if ok {
		return
	}

	visited[coordinates] = struct{}{}

	dfs(board, i, j+1, visited)
	dfs(board, i+1, j, visited)
	dfs(board, i, j-1, visited)
	dfs(board, i-1, j, visited)
}
