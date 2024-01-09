func maxSlidingWindow(nums []int, k int) []int {
	var (
		ans   []int
		deque []int
	)

	deque = append(deque, nums[0])
	for i := 1; i < k; i++ {
		deque = parse(deque, nums[i])
	}

	ans = append(ans, deque[0])
	for i := k; i < len(nums); i++ {
		if deque[0] == nums[i-k] {
			deque = deque[1:]
			if len(deque) == 0 {
				deque = append(deque, nums[i])
			}
		}
		deque = parse(deque, nums[i])
		ans = append(ans, deque[0])
	}

	return ans
}

func parse(deque []int, num int) []int {
	if deque[0] < num {
		deque = []int{num}
	} else if deque[len(deque)-1] > num {
		deque = append(deque, num)
	} else {
		var j = len(deque) - 1
		for deque[j] < num {
			j--
		}

		deque = append(deque[:j+1], num)
	}

	return deque
}