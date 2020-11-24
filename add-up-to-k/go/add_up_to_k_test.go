package adduptok

import "testing"

func TestIsAddUp(t *testing.T) {
	for _, test := range addUpToKCases {
		actual := IsAddUp(test.list, test.k)
		if actual != test.expected {
			t.Fatalf("FAIL : IsAddUp(%v, %d), expected [%t], got [%t] \n", test.list, test.k, test.expected, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkNearestLarger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range addUpToKCases {
			IsAddUp(test.list, test.k)
		}
	}
}
