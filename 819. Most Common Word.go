func mostCommonWord(paragraph string, banned []string) string {
	ban := make(map[string]int)
	para := make(map[string]int)
	str := ""

	for _, val := range banned {
		ban[val] = 1
	}

	for _, val2 := range paragraph {
		if val2 == ' ' || val2 == '!' || val2 == '?' || val2 == '\'' || val2 == ';' || val2 == ',' || val2 == '.' {
			para[str]++
			str = ""
		} else {
			str = str + strings.ToLower(string(val2))
		}
	}
	para[str]++
	count := 0
	result := ""

	for key, value := range para {
		if key != "" && value > count && ban[key] == 0 {
			count = value
			result = key
		}
	}
	return result
}    