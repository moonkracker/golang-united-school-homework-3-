package homework

func average(input [15]float32) (result float32) {
	for _, value := range input {
		result += value
	}
	return result / 15
}
