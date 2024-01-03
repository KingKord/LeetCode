
func maxArea(height []int) int {
	var (
		L   = 0
		R   = len(height) - 1
		max = 0
	)

	for L <= R {
		if height[L] < height[R] {
			max = Max(max, (R-L)*height[L])
			L++
		} else {
			max = Max(max, (R-L)*height[R])
			R--
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}