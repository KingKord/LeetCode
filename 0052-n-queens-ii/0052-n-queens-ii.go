func totalNQueens(n int) int {
	return recursion(nil, n, 0, 0)
}

func recursion(used []int, n, i, answer int) int {
	if i == n {
		return answer + 1
	}
	for j := 0; j < n; j++ {
		if isValid(used, j, i) {
			answer = recursion(append(used, j), n, i+1, answer)
		}
	}

	return answer
}

func isValid(used []int, x, y int) bool {
	for i, j := range used {
		if j == x || y-i == x-j || y-i == j-x {
			return false
		}
	}

	return true
}