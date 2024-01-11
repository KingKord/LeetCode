func partition(s string) [][]string {
	var ans [][]string
	ans = recursion(s, nil, ans)
	return ans
}

func recursion(s string, now []string, ans [][]string) [][]string {
	if len(s) == 0 {
		copied := make([]string, len(now))
		copy(copied, now)
		return append(ans, copied)
	}
	for i := 0; i < len(s); i++ {
		str := s[:i+1]
		if isPalindrome(str) {
			ans = recursion(s[i+1:], append(now, str), ans)
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	var j = len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
