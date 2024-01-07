func maxSubArray(nums []int) int {
	var (
		sum = 0
		max = -10000
	)

	for _, num := range nums {
		sum = Max(num, sum+num)
		max = Max(max, sum)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}