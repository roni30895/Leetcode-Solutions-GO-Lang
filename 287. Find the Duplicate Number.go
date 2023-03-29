func findDuplicate(nums []int) int {
	l := make(map[int]int)
	for key, val := range nums {
		if _, exists := l[val]; exists {
			return val
		}
		l[val] = key
	}
	return 0
}