package productofotherlements

// ProductOf returns a new array such that each element at index i of the new array
// is the product of all the numbers in the original array except the one at i.
func ProductOf(list []int) []int {
	total := 1
	result := make([]int, len(list))
	for _, v := range list {
		total = total * v
	}

	for i, v := range list {
		result[i] = total / v
	}

	return result
}
