func findMin(nums []int) int {
	var (
		L = 0
		R = len(nums) - 1
	)

	for L < R {
		m := (L + R) / 2

		if nums[L] < nums[R] {
			return nums[L]
		}

		if nums[L] <= nums[m] {
			L = m + 1
		} else {
			R = m
		}
	}

	return  nums[L]
}
