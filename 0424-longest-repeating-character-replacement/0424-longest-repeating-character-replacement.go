func characterReplacement(s string, k int) int {
	var (
		hashmap = make(map[uint8]int)
		max     = -k + 1
		now     = s[0]
		count   = 0
		last    = 0
	)
	hashmap[s[0]]++
	for i := 1; i < len(s); i++ {
		char := s[i]
		hashmap[char]++
		if char != now {
			count++
			for count > k {
				for s[last] == now {
					hashmap[s[last]]--
					last++
				}
				now = s[last]
				count = i - hashmap[s[last]] - last + 1
			}
		}

		max = Max(max, i-last+1-count)
	}
	last++
	for ; last < len(s); last++ {
		char := s[last]
		if char != now {
			hashmap[now]--
			for s[last] == now {
				hashmap[s[last]]--
				last++
			}

			count = len(s) - hashmap[s[last]] - last + 1
		}
		
		max = Max(max, len(s)-last+1-count)
	}
	if max+k > len(s) {
		return len(s)
	}
	return max + k
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
