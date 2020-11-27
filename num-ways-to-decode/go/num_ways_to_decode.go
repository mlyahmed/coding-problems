package numwaystodecode

import (
	"strconv"
)

// NumWaysToDecode returns how many ways the given code can be decoded.
func NumWaysToDecode(code string) int {
	if len(code) <= 1 {
		return 1
	}

	if v, _ := strconv.Atoi(code[0:2]); v > 26 {
		return NumWaysToDecode(code[1:])
	}

	return NumWaysToDecode(code[1:]) + NumWaysToDecode(code[2:])
}
