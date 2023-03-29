func isSameAfterReversals(num int) bool {
	rev_num1 := rev(num)
	rev_num2 := rev(rev_num1)
	if rev_num2 == num {
		return true
	}
	return false
}
func rev(num int) int {
	rev := 0
	for num > 0 {
		rev = (10 * (rev)) + num%10
		num = num / 10
	}
	return rev
}