func longestConsecutive(nums []int) int {
	h := make(map[int]struct{})
	max := 0

	for _, num := range nums {
		h[num] = struct{}{}
	}

	length := 1

	for key := range h {
		_, ok := h[key - 1]
		if ok {
			continue
		}
		for i := key + 1; ; i++ {
			_, ok := h[i]
			if !ok {
				break
			}
			length++
		}

		max = Max(length, max)
		length = 1
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}