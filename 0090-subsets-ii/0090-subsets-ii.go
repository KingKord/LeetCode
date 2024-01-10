func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	ans = recursion(nums, nil, 0, ans)

	return ans
}

func recursion(nums, input []int, depth int, ans [][]int) [][]int {
	if depth == len(nums) {

		for i := 0; i < len(ans); i++ {
			if len(ans[i]) == len(input) && slices.Equal(ans[i], input) {
				return ans
			}
		}

		copied := make([]int, len(input))
		copy(copied, input)
		return append(ans, copied)
	}

	ans = recursion(nums, input, depth+1, ans)
	ans = recursion(nums, append(input, nums[depth]), depth+1, ans)
	return ans
}