package homework

func Average(input [15]float32) float32 {
	return average(input)
}

func average(input [15]float32) (result float32) {
	//Place your code here
	result = 0.0
	for _, e := range input {
		result += e
	}
	return result / float32(len(input))
}
