package firstmissinginteger

import (
	"testing"
)

func TestFirstMissingInteger(t *testing.T) {
	for _, test := range testCases {
		actual := FirstMissingInteger(test.list)
		if test.expected != actual {
			t.Fatalf("Fail: FirstMissingInteger(%v). Expected %v, got %v", test.list, test.expected, actual)
		}
		t.Logf("PASS : %v", test.description)
	}
}

func BenchmarkFirstMissingInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			FirstMissingInteger(test.list)
		}
	}
}
