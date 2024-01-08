func longestOnes(nums []int, k int) int {
	var (
		Left      = 0
		flipCount = 0
		max       = 0
		i         = 0
	)

	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			flipCount++
			if flipCount > k {
				for nums[Left] == 1 {
					Left++
				}
				Left++
				flipCount--
			}
		}
		max = Max(max, i-Left+1)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}