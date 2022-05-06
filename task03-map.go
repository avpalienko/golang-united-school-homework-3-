package homework

import "sort"

func SortMapValues(input map[int]string) []string {
	return sortMapValues(input)
}
func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	result = make([]string, len(input))
	for i, k := range keys {
		result[i] = input[k]
	}
	return result
}
