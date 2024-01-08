func contains(now, required map[uint8]int) bool {
	for key, val := range required {
		if now[key] < val {
			return false
		}
	}

	return true
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	if len(t) == 1 {
		for i := 0; i < len(s); i++ {
			if t[0] == s[i] {
				return string(t[0])
			}
		}
		return ""
	}

	var (
		Left     = 0
		Right    = 0
		maxLeft  = 0
		maxRight = 0
		min      = math.MaxInt
		required = make(map[uint8]int)
		now      = make(map[uint8]int)
	)

	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}

	if len(t) == len(s) {
		for i := 0; i < len(s); i++ {
			now[s[i]]++
		}
		if contains(now, required) {
			return s
		} else {
			return ""
		}
	}
	for Left != len(s)-1 {
		for ; Right != len(s) && !contains(now, required); Right++ {
			now[s[Right]]++
		}
		if Right-Left < min && contains(now, required) {
			min = Right - Left
			maxLeft = Left
			maxRight = Right
		}

		now[s[Left]]--
		Left++
		for count := 0; Left < Right && count == 0; Left++ {
			count = required[s[Left]]
			now[s[Left]]--
		}

		Left--
		now[s[Left]]++
	}

	if Right-Left < min && contains(now, required) {
		return s[Left:Right]
	}
	if Left-1 >= 0 {
		Left--
		now[s[Left]]++
		if Right-Left < min && contains(now, required) {
			return s[Left:Right]
		}
	}
	return s[maxLeft:maxRight]
}