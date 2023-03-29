func heightChecker(heights []int) int {
	count := 0
	hts := []int{}
	for _, value := range heights {
		hts = append(hts, value)
	}
	sort.Ints(heights)
	for i := 0; i < len(heights); i++ {
		if hts[i] != heights[i] {
			count++
		}
	}
	return count
}