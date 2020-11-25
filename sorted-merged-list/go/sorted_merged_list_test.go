package sortedmergedlist

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	for _, test := range sortedMergedListCases {
		actual := Merge(test.list)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("Fail: Merge(%v). Expected %v, got %v", test.list, test.expected, actual)
		}
		t.Logf("PASS : %v", test.description)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortedMergedListCases {
			Merge(test.list)
		}
	}
}
