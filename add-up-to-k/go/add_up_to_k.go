package adduptok

// IsAddUp returns whether any two numbers from the list add up to k.
// Time Complexity : O(n), n = len(list).
func IsAddUp(list []int, k int) bool {
	index := make(map[int]int, len(list))
	for _, v := range list {
		if index[v] != 0 {
			return true
		} else if v < k {
			index[k-v] = v
		}
	}
	return false
}
