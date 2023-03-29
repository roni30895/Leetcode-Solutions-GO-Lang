func vowelStrings(words []string, left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		word := words[i]
		length := len(word)
		if length == 1 && isin(word[0]) {
			count++
		} else if length == 2 && isin(word[0]) && isin(word[1]) {
			count++
		} else if isin(word[0]) && isin(word[length-1]) {
			count++
		}
	}
	return count

}

func isin(char byte) bool {
	if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
		return true
	}
	return false
}