package numwaystodecode

import (
	"testing"
)

func TestNumWaysToDecode(t *testing.T) {
	for _, test := range testCases {
		actual := NumWaysToDecode(test.code)
		if test.expected != actual {
			t.Fatalf("Fail: NumWaysToDecode(%v). Expected %v, got %v", test.code, test.expected, actual)
		}
		t.Logf("PASS : %v", test.description)
	}
}

func BenchmarkNumWaysToDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			NumWaysToDecode(test.code)
		}
	}
}
