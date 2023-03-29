func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i != nums[i] {
			return i
		}
	}
	return nums[n-1] + 1
}