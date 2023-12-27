
func merge(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		intervalI := intervals[i]
		var interval = intervalI
		for j := i + 1; j < len(intervals); j++ {
			intervalJ := intervals[j]
			startJ := intervalJ[0]
			endJ := intervalJ[1]

			if interval[0] == startJ {
				interval = []int{startJ, Max(endJ, interval[1])}
				i++
			} else if startJ <= interval[1] && endJ > interval[1] {
				interval[1] = endJ
				i++
			} else if startJ <= interval[1] {
				i++
			}

			if startJ > interval[1] {
				i = j - 1
				break
			}
		}

		res = append(res, interval)
	}

	return res
}


func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}