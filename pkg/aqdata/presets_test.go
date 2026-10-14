package aqdata

import (
	"testing"
)

// checks the preset list methods
func checkPresetList[T presetLike](t *testing.T, list *presetList[T], first T) {
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
	// check the FromName method
	got, err = list.FromName(first.ToString())
	if err != nil {
		t.Fatal(err)
	}
	if got != first {
		t.Errorf("FromName(%q) = %v, expected %v", first.ToString(), got, first)
	}
	// check the FromName method with missing preset
	if _, err := list.FromName("missing"); err == nil {
		t.Error("FromName(missing) succeeded, expected error")
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

func checkPreset[T presetLike](t *testing.T, preset T, name string) {
	t.Helper()
	// check the ToString method
	if got := preset.ToString(); got != name {
		t.Errorf("ToString() = %q, expected %q", got, name)
	}
	// check the Index method
	if i, ok := preset.Index(); !ok || i != 0 {
		t.Errorf("Index() = %d, %v, expected 0, true", i, ok)
	}
	// check the Index method with zero value
	var zero T
	if _, ok := zero.Index(); ok {
		t.Error("Index() of zero value succeeded, expected false")
	}
}

// tests the preset list methods
func TestPresetMethods(t *testing.T) {
	t.Parallel()
	checkPreset(t, Aq1F1, "f1")
	checkPresetList(t, &aq1Presets, Aq1F1)
	checkPreset(t, Aq2AqYukkuri, "aq_yukkuri")
	checkPresetList(t, &aq2Presets, Aq2AqYukkuri)
	checkPreset(t, Aq10Custom, "custom")
	checkPresetList(t, &aq10Presets, Aq10Custom)
}

// duplicate define method should panic
func TestDefineDuplicatePanics(t *testing.T) {
	t.Parallel()
	var list presetList[Aq1Preset]
	list.define("once")
	defer func() {
		if recover() == nil {
			t.Error("duplicate define succeeded, expected panic")
		}
	}()
	list.define("once")
}
