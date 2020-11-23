package nearestlarger

// NearestLarger returns the index of the nearest larger number of the number at index i.
// Time Complexity : O(n), n = len(elements).
func NearestLarger(elements []int, i int) int {
	if i >= len(elements) {
		return -1
	}

	for x, xnext, xprevious := 1, true, true; xnext || xprevious; x++ {

		nearest := i
		xnext, xprevious = (i+x < len(elements)), (i-x >= 0)

		if xnext && elements[i+x] > elements[nearest] {
			nearest = i + x
		}

		if xprevious && elements[i-x] > elements[nearest] {
			nearest = i - x
		}

		if nearest != i {
			return nearest
		}

	}

	return -1
}
