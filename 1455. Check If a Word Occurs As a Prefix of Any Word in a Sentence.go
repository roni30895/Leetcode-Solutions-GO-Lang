func isPrefixOfWord(sentence string, searchWord string) int {
	wl := len(searchWord)
	sent := strings.Split(sentence, " ")
	for key, val := range sent {
		word := string(val)
		if len(word) >= wl && word[:wl] == string(searchWord) {
			return key + 1
		}
	}
	return -1
}