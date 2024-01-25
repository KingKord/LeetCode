func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		vertexes = make([][]int, numCourses)
		colors   = make([]int, numCourses)
		ans      []int
	)

	for _, prerequisite := range prerequisites {
		to, from := prerequisite[0], prerequisite[1]
		vertexes[from] = append(vertexes[from], to)
	}

	for i := range vertexes {
		if colors[i] == 0 {
			ans = dfs(vertexes, i, colors, ans)
		}
	}

	if len(ans) == numCourses {
		slices.Reverse(ans)
		return ans
	}
	return nil
}

func dfs(vertexes [][]int, i int, colors []int, ans []int) []int {
	colors[i] = 1

	for _, neighbor := range vertexes[i] {
		if colors[neighbor] == 1 {
			return nil
		} else if colors[neighbor] == 0 {
			ans = dfs(vertexes, neighbor, colors, ans)
		}
	}

	colors[i] = 2
	return append(ans, i)
}
