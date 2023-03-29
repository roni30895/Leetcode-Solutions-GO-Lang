func searchInsert(nums []int, target int) int {
	for key, value := range nums {
		if target == value {
			return key
		}
	}
	nums = append(nums, target)
	sorted_nums := sort(nums)
	return searchInsert(sorted_nums, target)

}
func sort(nums []int) []int {

	for j := len(nums) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}