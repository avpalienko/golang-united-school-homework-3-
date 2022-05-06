package homework

func Reverse(input []int64) []int64 {
	return reverse(input)
}

func reverse(input []int64) (result []int64) {
	//Place your code here
	inputLen := len(input)
	result = make([]int64, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		result[j] = n
	}
	return result
}
