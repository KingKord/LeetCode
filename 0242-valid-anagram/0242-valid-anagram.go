func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		count := hashMap[t[i]]
		if count == 0 {
			return false
		}
		hashMap[t[i]]--
	}

	return true
}