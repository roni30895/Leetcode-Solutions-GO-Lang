func countWords(words1 []string, words2 []string) int {
	L1 := make(map[string]int)
	L2 := make(map[string]int)
	result := 0
	for _, val1 := range words1 {
		L1[val1] = L1[val1] + 1
	}
	for _, val2 := range words2 {
		L2[val2] = L2[val2] + 1
	}
	for key, value := range L1 {
		if index, exists := L2[key]; exists {
			if index == 1 && value == 1 {
				result++
			}
		}
	}
	return result
}