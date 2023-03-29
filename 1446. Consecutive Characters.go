func maxPower(s string) int {
	if s == "" {
		return 0
	}
	count := 1
	result := 1
	left := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == left {
			count++
			if count > result {
				result = count
			}
		} else {
			count = 1
			left = s[i]
		}
	}
	return result
}


 
     