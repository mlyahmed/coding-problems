package nearestLarger

import "testing"

func TestNearestLarger(t *testing.T) {
  for _, test := range nearestLargerCases {

    actual := NearestLarger(test.input, test.element)

    if actual != test.expected {
      t.Fatalf("FAIL : NearestLarger(%v, %d), expected %d, got %d \n", test.input, test.element, test.expected, actual)
    }
    t.Logf("PASS: %s", test.description)

  }
}

func BenchmarkNearestLarger(b *testing.B) {
  for i:=0; i < b.N; i++ {
    for _, test := range nearestLargerCases {
      NearestLarger(test.input, test.element)
    }
  }
}
