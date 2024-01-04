
func search(nums []int, target int) int {
	var (
		L = 0
		R = len(nums) - 1
	)

	for L < R {
		m := (L + R) / 2
		if nums[m] < target {
			L = m + 1
		} else {
			R = m
		}
	}

	if nums[L] == target {
		return L
	}
	return -1
}
