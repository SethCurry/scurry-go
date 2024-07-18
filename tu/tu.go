package tu

import "testing"

// SliceEquality is a testing helper that checks that two slices are equal.
// It will report an error if the slices have different lengths, or if the
// values at any indexes do not match.
func SliceEquality[T comparable](t *testing.T, expected []T, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("expected slice of length %d, but got a slice of length %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("value at index %d does not match; expected %v but got %v", i, expected[i], actual[i])
		}
	}
}
