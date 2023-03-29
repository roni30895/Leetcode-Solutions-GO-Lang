func firstUniqChar(s string) int {
	L := make(map[string]int)
	for _, val := range s {
		L[string(val)] = L[string(val)] + 1

	}

	for key, val1 := range s {
		if index, exists := L[string(val1)]; exists {
			if index == 1 {
				return key
			}
		}
	}
	return -1
}