import "fmt"

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	L := make(map[int]int)
	result := []int{}
	for j := 1; j < length+1; j++ {
		result = append(result, j)
	}
	output := []int{}
	for i := 0; i < length; i++ {
		L[nums[i]] = 0
	}

	for _, v := range result {
		if _, ok := L[v]; !ok {
			output = append(output, v)
		}
	}

	return output
}