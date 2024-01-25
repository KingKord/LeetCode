func minimumLength(s string) int {
	var (
		left        = 0
		prefixStart = 0
		right       = len(s) - 1
		suffixStart = len(s) - 1
	)

	for left+1 < right && s[left] == s[left+1] {
		left++
	}
	for left < right-1 && s[right-1] == s[right] {
		right--
	}

	for s[left] == s[right] && left < right {
		left++
		prefixStart = left
		right--
		suffixStart = right
		for left+1 < right && s[left] == s[left+1] {
			left++
		}
		for left < right-1 && s[right-1] == s[right] {
			right--
		}
	}

	return suffixStart - prefixStart + 1
}