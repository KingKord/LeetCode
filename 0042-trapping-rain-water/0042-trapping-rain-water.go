func trap(height []int) int {
	var (
		max = 0
		sum = 0
	)

	for i := 0; i < len(height); i++ {
		if height[max] < height[i] {
			max = i
		}
	}

	lastMax := 0
	for i := 0; i <= max; i++ {
		if lastMax < height[i] {
			lastMax = height[i]
		} else {
			sum += lastMax - height[i]
		}
	}
	lastMax = 0

	for i := len(height) - 1; i >= max; i-- {
		if lastMax < height[i] {
			lastMax = height[i]
		} else {
			sum += lastMax - height[i]
		}
	}

	return sum
}