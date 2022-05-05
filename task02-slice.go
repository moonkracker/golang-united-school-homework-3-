package homework



func reverse(input []int64) (result []int64) {
	for _, value := range input {
		result = append([]int64{value}, result...)
	}
	return result
}
