
func isValidSudoku(board [][]byte) bool {
	boxes := make([]map[byte]struct{}, 3)
	var columns []map[byte]struct{}
	for i := 0; i < 9; i++ {
		columns = append(columns, make(map[byte]struct{}))
	}

	for i := 0; i < len(board); i++ {
		row := make(map[byte]struct{})
        if i % 3 == 0 {
            for j := 0; j < 3; j++ {
                boxes[j] = make(map[byte]struct{})
            }
        }

		for j := 0; j < len(board[i]); j++ {
			now := board[i][j]
			if now == '.' {
				continue
			}
			_, ok := row[now]
			if ok {
				fmt.Println(i, j, "row", now)
				return false
			}

			row[now] = struct{}{}

			_, ok = columns[j][now]
			if ok {
				fmt.Println(i, j, "columns")
				return false
			}
			columns[j][now] = struct{}{}

			_, ok = boxes[j/3][now]
			if ok {
				fmt.Println(i, j, "boxes")
				return false
			}
			boxes[j/3][now] = struct{}{}
		}
        for j := 0; j < 3; j++ {
			fmt.Print(boxes[j]," ")
		}
        fmt.Println()

	}

	return true
}