
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int
	var (
		first  = 0
		second = 0
		left   = 0
		right  = 0
	)

	for first < len(firstList) && second < len(secondList) {
		left = Max(firstList[first][0], secondList[second][0])
		right = Min(firstList[first][1], secondList[second][1])
		if left <= right {
			ans = append(ans, []int{left, right})
		}
		if secondList[second][1] < firstList[first][1] {
			second++
		} else {
			first++
		}
	}
	return ans
}
