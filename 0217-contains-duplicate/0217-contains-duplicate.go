func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		_, ok := hashMap[nums[i]]
		if ok {
			return true
		}
		hashMap[nums[i]] = struct{}{}
	}
	return false
}