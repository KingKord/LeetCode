func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	ans = recursion(candidates, nil, target, 0, 0, ans)

	return ans
}

func recursion(candidates, res []int, target, sum, start int, ans [][]int) [][]int {
	prev := -1
	for i := start; i < len(candidates); i++ {
		value := candidates[i]
		if value == prev {
			continue
		}
		prev = value
		sumNow := sum + value

		if sumNow < target {
			ans = recursion(candidates, append(res, value), target, sumNow, i+1, ans)
		} else {
			if sumNow == target {
				res = append(res, value)
				for _, arr := range ans {
					if len(arr) == len(res) && slices.Equal(arr, res) {
						return ans
					}
				}
				copied := make([]int, len(res))
				copy(copied, res)
				return append(ans, copied)
			}
			return ans
		}
	}

	return ans
}
