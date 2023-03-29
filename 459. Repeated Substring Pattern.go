func repeatedSubstringPattern(s string) bool {

	for i := 1; i <= len(s)/2; i++ {
		if (strings.Repeat(string(s[:i]), len(s)/i) == s) && len(s)%i == 0 {
			return true
		}
	}
	return false
}

