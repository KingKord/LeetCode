func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		vertexes = make([][]int, numCourses)
		colors   = make([]int, numCourses)
	)

	for _, prerequisite := range prerequisites {
		to, from := prerequisite[0], prerequisite[1]
		vertexes[from] = append(vertexes[from], to)
	}

	for i := range vertexes {
		if colors[i] == 0 && !dfs(vertexes, i, colors) {
			return false
		}
	}

	return true
}

func dfs(vertexes [][]int, i int, colors []int) bool {
	colors[i] = 1

	for _, neighbor := range vertexes[i] {
		if colors[neighbor] == 1 {
			return false
		} else if colors[neighbor] == 0 {
			if !dfs(vertexes, neighbor, colors) {
				return false
			}
		}
	}
	colors[i] = 2

	return true
}