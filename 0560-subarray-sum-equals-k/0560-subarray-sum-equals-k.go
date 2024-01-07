func subarraySum(nums []int, k int) int {
	var (
		counter  int
		sum      int
		prefixes = make(map[int]int)
	)

	prefixes[0] = 1
	for _, num := range nums {
		sum += num

		counter += prefixes[sum-k]
		prefixes[sum]++

	}

	return counter
}