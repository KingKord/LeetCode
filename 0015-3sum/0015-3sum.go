func threeSum(nums []int) [][]int {
	var (
		res [][]int
		L   int
		R   = 1
		C   = -1
	)
	sort.Ints(nums)
	for {
		C++
		L = C + 1
		R = len(nums) - 1
		if L >= R {
			return res
		}
		if C > 0 && nums[C] == nums[C-1] {
			continue
		}
		for L < R {
			sum := nums[C] + nums[L] + nums[R]
			if sum > 0 {
				R--
			} else if sum < 0 {
				L++
			} else {
				res = append(res, []int{nums[C], nums[L], nums[R]})
				L++
				R--

				for L < R && nums[L-1] == nums[L] {
					L++
				}

				for L < R && nums[R] == nums[R+1] {
					R--
				}
			}
		}
	}

}