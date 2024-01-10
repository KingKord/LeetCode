func subsets(nums []int) [][]int {
	var ans [][]int
	ans = find(nums, []int{}, ans, 0)

	return ans
}

func find(nums, res []int, ans [][]int, depth int) [][]int {
	if depth == len(nums) {
		toCopy := make([]int, len(res))
		copy(toCopy, res)
		return append(ans, toCopy)
	}
	ans = find(nums, res, ans, depth+1)
	ans = find(nums, append(res, nums[depth]), ans, depth+1)
	return ans
}