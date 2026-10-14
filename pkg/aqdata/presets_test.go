package aqdata

import (
	"testing"
)

func checkPresets[T comparable](t *testing.T, list *presetList[T], first T) {
	t.Helper()
	got, err := list.fromIndex(0)
	if err != nil {
		t.Fatal(err)
	}
	if got != first {
		t.Errorf("fromIndex(0) = %v, expected %v", got, first)
	}
	for _, i := range []int{-1, list.length(), 114514} {
		if _, err := list.fromIndex(i); err == nil {
			t.Errorf("fromIndex(%d) succeeded, expected error", i)
		}
	}
	n := list.length()
	for _, tc := range []struct{ limit, expected int }{
		{-1, 0},
		{5, min(5, n)},
		{n, n},
		{114514, n},
	} {
		if got := len(list.getPresets(tc.limit)); got != tc.expected {
			t.Errorf("getPresets(%d) len = %d, expected %d", tc.limit, got, tc.expected)
		}
	}
}

func TestPresets(t *testing.T) {
	t.Parallel()
	checkPresets(t, &aq1Presets, Aq1F1)
	checkPresets(t, &aq2Presets, Aq2AqYukkuri)
	checkPresets(t, &aq10Presets, Aq10Custom)
}
