package aqdata

import (
	"testing"
)

// TestAq1Presets tests presets for aquestalk 1
func TestAq1Presets(t *testing.T) {
	t.Parallel()
	result, err := ToAq1Preset(0)
	if err != nil {
		t.Error(err)
	}
	if result != F1aq1 {
		t.Errorf(
			"expected %d, got %d",
			result,
			F1aq1,
		)
	}
	allAqs, err := GetAq1Presets(
		-1,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allAqs) != 0 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			0,
			len(allAqs),
		)
	}
	allAqs, err = GetAq1Presets(
		114514,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allAqs) != 9 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			9,
			len(allAqs),
		)
	}
	allAqs, err = GetAq1Presets(
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
