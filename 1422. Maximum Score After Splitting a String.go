func maxScore(s string) int {
	var result = 0
	var sum = 0
	for i := 0; i < len(s)-1; i++ {
		zeros := strings.Count(s[:i+1], "0")
		ones := strings.Count(s[i+1:], "1")
		sum = zeros + ones
		if sum > result {
			result = sum
		}
	}
	return result
}