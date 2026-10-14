package aqdata

import (
	"testing"
)

// TestAq1Presets tests presets for aquestalk 1
func TestAq1Presets(t *testing.T) {
	t.Parallel()
	result, err := aq1Presets.fromIndex(0)
	if err != nil {
		t.Error(err)
	} else if result != Aq1F1 {
		t.Errorf("expected %s, got %s", Aq1F1.ToString(), result.ToString())
	}
	result, err = aq1Presets.fromIndex(114514)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
		)
	}
	result, err = aq1Presets.fromIndex(-114514)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
		)
	}
	result, err = aq1Presets.fromIndex(9)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
		)
	}
	allAqs := aq1Presets.getPresets(
		-1,
	)
	if len(allAqs) != 0 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			0,
			len(allAqs),
		)
	}
	allAqs = aq1Presets.getPresets(
		114514,
	)
	if len(allAqs) != 9 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			9,
			len(allAqs),
		)
	}
	allAqs = aq1Presets.getPresets(
		5,
	)
	if len(allAqs) != 5 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			5,
			len(allAqs),
		)
	}
}

// TestAq2Presets tests functions related to presets of aquestalk2
func TestAq2Presets(t *testing.T) {
	t.Parallel()
	result0, err := aq2.fromIndex(0)
	if err != nil {
		t.Error(err)
	}
	if result0 != Aq2AqYukkuri {
		t.Errorf(
			"expected %s, got %s",
			result0.ToString(),
			Aq2AqYukkuri.ToString(),
		)
	}
	_, err = aq2.fromIndex(114514)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	_, err = aq2.fromIndex(-1)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	allAqs := aq2.getPresets(
		-1,
	)
	if len(allAqs) != 0 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			0,
			len(allAqs),
		)
	}
	allAqs = aq2.getPresets(
		114514,
	)
	if len(allAqs) != 16 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			16,
			len(allAqs),
		)
	}
	allAqs = aq2.getPresets(
		16,
	)
	if len(allAqs) != 16 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			16,
			len(allAqs),
		)
	}
}

// TestAq10Presets tests presets of aquestalk10
func TestAq10Presets(t *testing.T) {
	t.Parallel()
	result0, err := aq10.fromIndex(0)
	if err != nil {
		t.Error(err)
	}
	if result0 != Aq10Custom {
		t.Errorf(
			"expected %s, got %s",
			result0.ToString(),
			Aq10Custom.ToString(),
		)
	}
	_, err = aq10.fromIndex(114514)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	_, err = aq10.fromIndex(-1)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	allAqs := aq10.getPresets(
		-1,
	)
	if len(allAqs) != 0 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			0,
			len(allAqs),
		)
	}
	allAqs = aq10.getPresets(
		114514,
	)
	if len(allAqs) != 8 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			8,
			len(allAqs),
		)
	}
	allAqs = aq10.getPresets(
		8,
	)
	if len(allAqs) != 8 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			8,
			len(allAqs),
		)
	}
}
