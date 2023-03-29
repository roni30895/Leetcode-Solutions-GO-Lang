func findTheDifference(s string, t string) byte {
	var single byte = 0
	for j := 0; j < len(t); j++ {
		if len(s) > j {
			single = single ^ s[j]
		}
		single = single ^ t[j]
	}
	return single
}