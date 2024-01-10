func permute(nums []int) [][]int {
	var ans [][]int
	ans = recursion(nums,nil,ans)
	
	return ans
}

func recursion(left, used []int, ans [][]int) [][]int {
	if len(left) == 0 {
		copied := make([]int, len(used))
		copy(copied, used)
		ans = append(ans, copied)
		return ans
	}

	for i, num := range left {
		copied := make([]int, len(left))
		copy(copied, left)
		copied = append(copied[:i], copied[i+1:]...)
		ans = recursion(copied, append(used, num), ans)
	}
	
	return ans
}