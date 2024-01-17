type tuple struct {
	isPacific  bool
	isAtlantic bool
}

func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int
	var computed = make(map[[2]int]tuple)
	for i, row := range heights {
		for j, _ := range row {
			res := dfs(heights, computed, make(map[[2]int]struct{}), i, j, math.MaxInt, tuple{})
			if res.isPacific && res.isAtlantic {
				ans = append(ans, []int{i, j})
			}
            computed[[2]int{i, j}] = res

		}
	}
	return ans
}

func dfs(heights [][]int, computed map[[2]int]tuple, visited map[[2]int]struct{}, i, j, prev int, oceans tuple) tuple {
	if i < 0 || j < 0 {
		oceans.isPacific = true
		return oceans
	} else if i >= len(heights) || j >= len(heights[0]) {
		oceans.isAtlantic = true
		return oceans
	}

	if heights[i][j] > prev {
		return oceans
	}
	coordinates := [2]int{i, j}
	computedTuple, ok := computed[coordinates]
	if ok {
		return tuple{
			isAtlantic: computedTuple.isAtlantic || oceans.isAtlantic,
			isPacific: computedTuple.isPacific || oceans.isPacific,
		}
	}

	_, ok = visited[[2]int{i, j}]
	if ok {
		return oceans
	}
	visited[coordinates] = struct{}{}

	oceans = dfs(heights, computed, visited, i, j+1, heights[i][j], oceans)
	if oceans.isAtlantic && oceans.isPacific {
		return oceans
	}

	oceans = dfs(heights, computed, visited, i-1, j, heights[i][j], oceans)
	if oceans.isAtlantic && oceans.isPacific {
		return oceans
	}

	oceans = dfs(heights, computed, visited, i, j-1, heights[i][j], oceans)
	if oceans.isAtlantic && oceans.isPacific {
		return oceans
	}

	oceans = dfs(heights, computed, visited, i+1, j, heights[i][j], oceans)
	if oceans.isAtlantic && oceans.isPacific {
		return oceans
	}

	return oceans
}