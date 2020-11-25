package sortedmergedlist

// Merge returns a new sorted merged list from the sorted lists.
// Brute force version. Complexity : Time O(k*n pow 2), Space O( k*n)
func Merge(list [][]int) []int {
	result := []int{}

	for _, l := range list {
		result = append(result, l...)
	}

	for i := range result {
		for k := i + 1; k < len(result); k++ {
			if result[k] < result[i] {
				result[k], result[i] = result[i], result[k]
			}
		}
	}

	return result
}
