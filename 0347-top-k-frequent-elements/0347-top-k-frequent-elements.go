func topKFrequent(nums []int, k int) []int {
	var (
		hashmap = make(map[int]int)
		res     = make([]int, k)
	)

	for _, num := range nums {
		hashmap[num]++
	}

	for i := 0; i < k; i++ {
		count := 0
		el := 0

		for key, occurrence := range hashmap {
			if occurrence > count {
				el = key
				count = occurrence
			}
		}
		res[i] = el
		delete(hashmap, el)
	}

	return res
}