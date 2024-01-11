func solveNQueens(n int) [][]string {
	var ans [][]int
	ans = recursion(n, 0, nil, ans)

	var strAns [][]string
	for _, answer := range ans {
		arr := make([]string, n)
		for i, j := range answer {
			str := ""
			for k := 0; k < n; k++ {
				if k == j {
					str += "Q"
				} else {
					str += "."
				}
			}
			arr[i] = str
		}

		strAns = append(strAns, arr)
	}
	return strAns
}

func recursion(n int, i int, previous []int, ans [][]int) [][]int {
	if len(previous) == n {
		copied := make([]int, n)
		copy(copied, previous)
		return append(ans, copied)
	}

	for j := 0; j < n; j++ {
		if isValid(previous, j, i, n) {
			ans = recursion(n, i+1, append(previous, j), ans)
		}
	}
	return ans
}

func isValid(previous []int, x, y, n int) bool {
	for i, j := range previous {
		if j == x || y-i == j-x || y-i == x-j {
			return false
		}
	}

	return true
}
