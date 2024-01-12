func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) <= 2 {
		if len(nums) == 1 {
			return nums[0]
		} else if len(nums) == 2 {
			return max(nums[1], nums[0])
		}
	}
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}

	return max(dp[len(nums)-1], dp[len(nums)-2])
}