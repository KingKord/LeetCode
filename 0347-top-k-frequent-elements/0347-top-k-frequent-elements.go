func topKFrequent(nums []int, k int) []int {
	intMap := make(map[int]int)
	for _, num := range nums {
		intMap[num]++
	}

	var res []int
	for i := range intMap {
		res = append(res, i)
	}
	
	sort.Slice(res, func(i, j int) bool {
		return intMap[res[i]] > intMap[res[j]]
	})

	return res[:k]
}