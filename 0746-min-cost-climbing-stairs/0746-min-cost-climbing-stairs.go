func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2] + cost[i], dp[i-1] + cost[i])
	}
	
	return min(dp[len(cost)-1], dp[len(cost) -2])
}