package aqdata

import (
	"testing"
)

// checks the preset list methods
func checkPresets[T presetLike](t *testing.T, list *presetList[T], first T) {
	t.Helper()
	// check the FromIndex method
	got, err := list.FromIndex(0)
	if err != nil {
		t.Fatal(err)
	}
	if got != first {
		t.Errorf("FromIndex(0) = %v, expected %v", got, first)
	}
	// check the FromIndex method with out of range index
	for _, i := range []int{-1, list.Length(), 114514} {
		if _, err := list.FromIndex(i); err == nil {
			t.Errorf("FromIndex(%d) succeeded, expected error", i)
		}
	}
	// check the GetPresets method
	n := list.Length()
	for _, tc := range []struct{ limit, expected int }{
		{-1, 0},
		{5, min(5, n)},
		{n, n},
		{114514, n},
	} {
		if got := len(list.GetPresets(tc.limit)); got != tc.expected {
			t.Errorf("GetPresets(%d) len = %d, expected %d", tc.limit, got, tc.expected)
		}
	}
}

func TestPresets(t *testing.T) {
	t.Parallel()
	checkPresets(t, &aq1Presets, Aq1F1)
	checkPresets(t, &aq2Presets, Aq2AqYukkuri)
	checkPresets(t, &aq10Presets, Aq10Custom)
}
