func exist(board [][]byte, word string) bool {
	exclusions := make(map[[2]int]struct{})

	var dfs func(i, j, depth int) bool
	dfs = func(i, j, depth int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}

		pair := [2]int{i, j}
		_, ok := exclusions[pair]
		if ok {
			return false
		}
		if board[i][j] != word[depth] {
			return false
		}
		depth++
		if depth == len(word) {
			return true
		}

		exclusions[pair] = struct{}{}
		if dfs(i-1, j, depth) {
			return true
		}
		if dfs(i+1, j, depth) {
			return true
		}
		if dfs(i, j-1, depth) {
			return true
		}
		if dfs(i, j+1, depth) {
			return true
		}
		delete(exclusions, pair)
		return false
	}
	for i, bytes := range board {
		for j := range bytes {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}