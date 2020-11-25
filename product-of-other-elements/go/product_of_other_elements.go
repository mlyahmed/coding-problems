package productofotherlements

// ProductOf returns a new array such that each element at index i of the new array
// is the product of all the numbers in the original array except the one at i.
// Thanks to https://leetcode.com/problems/product-of-array-except-self/solution/.
// Time Comlexity : O(n).
// Space Complexity : O(n).
// where n=len(list).
func ProductOf(list []int) []int {
	result := make([]int, len(list))

	for i := range list {
		result[i] = 1
	}

	for i := 1; i < len(list); i++ {
		result[i] = result[i-1] * list[i-1]
	}

	for r, i := 1, len(list)-2; i >= 0; i-- {
		r = r * list[i+1]
		result[i] *= r
	}

	return result
}
