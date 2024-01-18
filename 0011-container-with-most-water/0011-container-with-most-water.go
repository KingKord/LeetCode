func maxArea(height []int) int {
	var (
		L   = 0
		R = len(height) - 1
		m = 0
	)

	for L <= R {
		if height[L] < height[R] {
			m = max(m, (R-L)*height[L])
			L++
		} else {
			m = max(m, (R-L)*height[R])
			R--
		}
	}

	return m
}