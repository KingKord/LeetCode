
func isPalindrome(s string) bool {
	var (
		left  = 0
		right = len(s) - 1
	)

	for left < right {
				if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right]))&& !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		} else {
			left++
            right--
		}
	}

	return true
}