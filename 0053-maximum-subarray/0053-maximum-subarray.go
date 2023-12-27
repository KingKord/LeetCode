func maxSubArray(nums []int) int {
	var (
		sum = nums[0]
		max = sum
	)

	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val > 0 {
			sum = Max(val, val+sum)
			max = Max(sum, max)
		} else {
			if sum < val {
				sum = val
			} else if sum+val > 0 {
				sum += val
			} else {
				max = Max(sum, max)
				sum = -100000
			}

			max = Max(sum, max)
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}