func isValid(s string) bool {
	var stack []uint8

	for i := 0; i < len(s); i++ {
		char := s[i]
		if len(stack) == 0 {
			if char == ')' || char == '}' || char == ']' {
				return false
			}
			stack = append(stack, char)
			continue
		}
		last := stack[len(stack)-1]
		switch char {
		case ')':
			if last != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if last != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if last != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, char)
		}
	}

	if len(stack) != 0 {
		return false
	}
    
	return true
}