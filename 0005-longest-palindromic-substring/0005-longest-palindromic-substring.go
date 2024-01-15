func longestPalindrome(s string) string {
	var ans = string(s[0])
	var countMax = 1

	s = "$" + strings.Join(strings.Split(s, ""), "#") + "*"

	for i := 1; i < len(s)-1; i++ {
		count := 0
		if s[i] != '#' {
			count = 1
		}
		left, right := i, i
		for j := 1; i+j < len(s) && i-j >= 0 && s[i-j] == s[i+j]; j++ {
			left--
			right++
			if s[i-j] != '#' {
				count += 2
			}
		}
		if countMax < count {
			countMax = count
			ans = s[left : right+1]
		}
	}

	return strings.Join(strings.Split(ans, "#"), "")
}
