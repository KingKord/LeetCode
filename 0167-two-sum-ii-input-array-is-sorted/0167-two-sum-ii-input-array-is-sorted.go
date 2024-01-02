func twoSum(numbers []int, target int) []int {
	var (
		L = 0
		R = len(numbers) - 1
	)

	for i := 0; i < len(numbers); i++ {
		sum := numbers[L] + numbers[R]
		if sum < target {
			L++
		} else if sum > target {
			R--
		} else {
			return []int{L + 1, R + 1}
		}
	}

	return nil
}

