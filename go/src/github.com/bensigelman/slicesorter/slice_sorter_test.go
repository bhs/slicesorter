package slicesorter

import "testing"

// Let's at least test something that sort.Ints() wouldn't handle on its own.
type NotAnInt int

func slicesAreEqual(t *testing.T, expected, actual []NotAnInt) {
	if len(actual) != len(expected) {
		t.Errorf("lens don't match (%v != %v)", len(actual), len(expected))
	}
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("%v != %v", actual[i], v)
		}
	}
}

func TestBasics(t *testing.T) {
	s := []NotAnInt{3, 4, 5, 8, 5, 2}
	expected := []NotAnInt{8, 5, 5, 4, 3, 2}
	SortSlice(s, func(l interface{}, r interface{}) bool { return l.(NotAnInt) > r.(NotAnInt) })
	slicesAreEqual(t, expected, s)

	expected2 := []NotAnInt{2, 3, 4, 5, 5, 8}
	SortSlice(s, func(l interface{}, r interface{}) bool { return l.(NotAnInt) < r.(NotAnInt) })
	slicesAreEqual(t, expected2, s)
}

func TestEdgeCases(t *testing.T) {
	s := []NotAnInt{}
	SortSlice(s, func(l interface{}, r interface{}) bool { return l.(NotAnInt) > r.(NotAnInt) })
	slicesAreEqual(t, []NotAnInt{}, s)
}
