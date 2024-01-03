func evalRPN(tokens []string) int {
	var stack []int

	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		if err != nil {
			res := 0
			switch tokens[i] {
			case "+":
				res = stack[len(stack)-1] + stack[len(stack)-2]
			case "-":
				res = stack[len(stack)-2] - stack[len(stack)-1]
			case "*":
				res = stack[len(stack)-2] * stack[len(stack)-1]
			case "/":
				res = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack[len(stack)-2] = res
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}