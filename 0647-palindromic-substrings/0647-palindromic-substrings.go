func countSubstrings(s string) int {
	var ans = len(s)

	s = "$" + strings.Join(strings.Split(s, ""), "#") + "*"

	for i := 1; i < len(s)-1; i++ {
		left, right := i, i

		for j := 1; i+j < len(s) && i-j >= 0 && s[i-j] == s[i+j]; j++ {
			left--
			right++
			if s[i-j] != '#' {
				ans++
			}
		}
	}

	return ans
}
