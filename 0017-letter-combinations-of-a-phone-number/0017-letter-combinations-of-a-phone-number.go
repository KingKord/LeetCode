func letterCombinations(digits string) []string {
    if digits == "" {
		return nil
	}
	var ans []string
	ans = recursion(digits, nil, ans)

	return ans
}

func recursion(digits string, now []int32, ans []string) []string {
	if digits == "" {
		return append(ans, string(now))
	}

	first := digits[0]
	digits = digits[1:]
	var arr []int32
	switch first {
	case '2':
		arr = []int32("abc")
	case '3':
		arr = []int32("def")
	case '4':
		arr = []int32("ghi")
	case '5':
		arr = []int32("jkl")
	case '6':
		arr = []int32("mno")
	case '7':
		arr = []int32("pqrs")
	case '8':
		arr = []int32("tuv")
	case '9':
		arr = []int32("wxyz")
	}

	for _, char := range arr {
		ans = recursion(digits, append(now, char), ans)
	}
	return ans
}