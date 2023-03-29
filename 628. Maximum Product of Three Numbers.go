func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return (nums[0] * nums[1] * nums[2])
	}
	sort.Ints(nums)
	i := len(nums)
	x := nums[i-1] * nums[i-2] * nums[i-3]
	y := nums[0] * nums[1] * nums[i-1]

	if x > y {
		return x
	}
	return y

}