func toLowerCase(s string) string {
	str := ""
	for _, value := range s {
		if 'A' <= value && value <= 'Z' {
			str = str + string(value+32)
		} else {
			str = str + string(value)
		}
	}
	return str
}