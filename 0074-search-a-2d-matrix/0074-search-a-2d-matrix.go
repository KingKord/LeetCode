func searchMatrix(matrix [][]int, target int) bool {
	var (
		rowLength = len(matrix[0])
		L         = 0
		R         = len(matrix[0])*len(matrix) - 1
	)

	for L < R {
		m := (L + R) / 2
		middleRow := m / rowLength
		middleCol := m % rowLength

		if matrix[middleRow][middleCol] == target {
			return true
		} else if matrix[middleRow][middleCol] > target {
			R = middleRow*rowLength + middleCol
		} else {
			L = middleRow*rowLength + middleCol + 1
		}
	}

	if matrix[L/rowLength][L%rowLength] == target {
		return true
	}
	return false
}