func binary(nums []int, target int) int {
	var (
		L = 0
		R = len(nums) - 1
	)

	for L < R {
		m := (L + R) / 2
		if nums[m] < target {
			R = m
		} else {
			L = m + 1
		}
	}

	return L
}
func dailyTemperatures(temperatures []int) []int {
	var (
		res   = make([]int, len(temperatures))
		stack []int
		idx   []int
	)

	for i, temperature := range temperatures {
		firstIdx := binary(stack, temperature)

		if stack != nil && temperature > stack[firstIdx] {
			for j := firstIdx; j < len(stack); j++ {
				res[idx[j]] = i - idx[j]
			}
			stack = stack[:firstIdx]
			idx = idx[:firstIdx]
		}

		stack = append(stack, temperature)
		idx = append(idx, i)
	}

	return res
}