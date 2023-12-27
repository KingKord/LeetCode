
func longestValidParentheses(s string) int {
	h := make(map[int]int)
	max := 0
	p := 0

	for _, char := range s { // 40 - "(", 41 - ")"
		if char == 40 {
			h[p]++
			p++
		} else if p == 0 {
			h = make(map[int]int)
			continue
		} else {
			p--
			count, ok := h[p+1]
			if !ok {
				h[p]++
			} else {
				delete(h, p+1)
				h[p] += count + 1
			}
			if max < h[p] {
				max = h[p]
			}
		}
	}

	return max
}
