package firstmissinginteger

// FirstMissingInteger returns the first positive integer in the list.
func FirstMissingInteger(list []int) int {

	for i := range list {
		for k := i + 1; k < len(list); k++ {
			if list[k] < list[i] {
				list[i], list[k] = list[k], list[i]
			}
		}

		if i >= 1 && list[i-1] >= 0 && (list[i]-list[i-1] > 1) {
			return list[i-1] + 1
		}
	}

	return list[len(list)-1] + 1
}
