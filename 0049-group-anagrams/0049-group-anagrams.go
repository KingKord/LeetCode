func groupAnagrams(strs []string) [][]string {
	var (
		ans     [][]string
		someMap = make(map[[26]int][]string)
	)


	for i := 0; i < len(strs); i++ {

		var alphabet [26]int
		str := strs[i]
		for j := 0; j < len(str); j++ {
			charNow := str[j] - 'a'
			alphabet[charNow]++
		}

		someMap[alphabet] = append(someMap[alphabet], str)
	}

	for _, strings := range someMap {
		ans = append(ans, strings)
	}
	return ans
}