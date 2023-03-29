func detectCapitalUse(word string) bool {
	count := 0
	for _, value := range word {
		if 'A' <= value && value <= 'Z' {
			count++
		}
	}
	if count == len(word) || count == 0 || (count == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}