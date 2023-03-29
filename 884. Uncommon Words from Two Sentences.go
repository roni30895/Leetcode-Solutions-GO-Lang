func uncommonFromSentences(s1 string, s2 string) []string {
	str1 := strings.Split(s1, " ")
	str2 := strings.Split(s2, " ")
	result := []string{}
	L := make(map[string]int)
	for _, value1 := range str1 {
		if _, ok := L[string(value1)]; ok {
			L[string(value1)]++
		} else {
			L[string(value1)] = 1
		}
	}
	for _, value2 := range str2 {
		if _, ok := L[string(value2)]; ok {
			L[string(value2)]++
		} else {
			L[string(value2)] = 1
		}
	}
	for key, val := range L {
		if val == 1 {
			result = append(result, key)
		}
	}
	return result
}