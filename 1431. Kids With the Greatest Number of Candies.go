func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	max := maximum(candies)
	for _, val := range candies {
		total := val + extraCandies
		if total >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
func maximum(arr []int) int {
	c := arr[0]
	for _, val := range arr {
		if c < val {
			c = val
		}
	}
	return c
}