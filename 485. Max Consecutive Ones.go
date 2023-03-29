func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > result {
				result = count
			}
		} else if nums[i] == 0 {
			count = 0
		}
	}
	return result
}