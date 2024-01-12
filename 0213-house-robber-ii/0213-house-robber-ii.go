func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(find(nums[:len(nums)-1]), find(nums[1:]))
}

func find(nums []int) int {
	if len(nums) <= 2 {
		if len(nums) == 1 {
			return nums[0]
		} else if len(nums) == 2 {
			return max(nums[1], nums[0])
		}
	}
	first := nums[0]
	second := nums[1]
	cur := nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		temp := max(first, second) + nums[i]
		first = second
		second = cur
		cur = temp
	}

	return max(cur, second)
}
