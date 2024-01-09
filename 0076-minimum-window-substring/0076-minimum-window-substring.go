func minWindow(s string, t string) string {
    if len(t) > len(s) {
		return ""
	}
	var (
		required  = make(map[uint8]int)
		window    = make(map[uint8]int)
		resLeft   = 0
		resRight  = -1
		maxLength = math.MaxInt
		need      = 0
		have      = 0
		last      = 0
	)

	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}
	need = len(required)

	for i := 0; i < len(s); i++ {
		char := s[i]
		_, ok := required[char]
		if ok {
			window[char]++
			count := window[char]
			if count == required[char] {
				have++
			}

			for need == have {
				lastChar := s[last]
				count, ok := window[lastChar]
				if ok {
					window[lastChar]--
					if count-1 < required[lastChar] {
						have--
						if i-last+1 < maxLength {
							resLeft = last
							resRight = i
							maxLength = i - last + 1
						}
					}
				}
				last++
			}
		}
	}

	return s[resLeft : resRight+1]
}
