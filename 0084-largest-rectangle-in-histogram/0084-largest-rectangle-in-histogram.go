
func binary(nums []int, target int) int {
	var (
		L = 0
		R = len(nums) - 1
	)

	for L < R {
		m := (L + R) / 2
		if nums[m] > target {
			R = m
		} else {
			L = m + 1
		}
	}

	return L
}

func largestRectangleArea(heights []int) int {
	var (
		max      int
		stack    []int
		idxStack []int
	)
	stack = append(stack, heights[0])
	idxStack = append(idxStack, 0)
	for i := 1; i < len(heights); i++ {
		height := heights[i]
		lastIdx := binary(stack, height)

		if stack[lastIdx] >= height {
			max = Max(max, (i-idxStack[lastIdx])*stack[lastIdx])
			stack[lastIdx] = height

			for j := lastIdx + 1; j < len(stack); j++ {
				max = Max(max, (i-idxStack[j])*stack[j])
			}
			stack = stack[:lastIdx+1]
			idxStack = idxStack[:lastIdx+1]
		}

		stack = append(stack, height)
		idxStack = append(idxStack, i)
	}

	for j := 0; j < len(stack); j++ {
		max = Max(max, (len(heights)-idxStack[j])*stack[j])
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
