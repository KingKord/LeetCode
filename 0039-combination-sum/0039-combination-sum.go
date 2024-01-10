func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	ans = recursion(candidates, nil, target, 0, 0, ans)

	return ans
}

func recursion(candidates, res []int, target, start, sum int, ans [][]int) [][]int {
	for i := start; i < len(candidates); i++ {
		value := candidates[i]
		newSum := sum + value
		if newSum < target {
			ans = recursion(candidates, append(res, value), target, i, newSum, ans)
		} else {
			if newSum == target {
				var copied = make([]int, len(res))
				copy(copied, res)
				ans = append(ans, append(copied, value))
			}
			break
		}
	}
	return ans
}