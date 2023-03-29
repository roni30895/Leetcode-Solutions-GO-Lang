func findSubarrays(nums []int) bool {
	L := make(map[int]bool)
	for k := 0; k < len(nums)-1; k++ {
		sum := nums[k] + nums[k+1]
		if _, exists := L[sum]; exists {
			return true
		}
		L[sum] = true
	}
	return false
}