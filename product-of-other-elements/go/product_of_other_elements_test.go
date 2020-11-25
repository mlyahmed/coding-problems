package productofotherlements

import (
	"reflect"
	"testing"
)

func TestProductOf(t *testing.T) {
	for _, test := range productOfOtherElementsCases {

		actual := ProductOf(test.list)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL : ProductOf(%v), expected %v, got %v \n", test.list, test.expected, actual)
		}
		t.Logf("PASS: %s", test.description)

	}
}

func BenchmarkProductOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range productOfOtherElementsCases {
			ProductOf(test.list)
		}
	}
}
