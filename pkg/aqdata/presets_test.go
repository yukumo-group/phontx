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
			"expected %s, got %s",
			result.ToString(),
			F1aq1.ToString(),
		)
	}
	result, err = ToAq1Preset(114514)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
		)
	}
	result, err = ToAq1Preset(-114514)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
		)
	}
	result, err = ToAq1Preset(9)
	if err == nil {
		t.Error(
			"expected to throw an error here!",
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

// TestAq2Presets tests functions related to presets of aquestalk2
func TestAq2Presets(t *testing.T) {
	t.Parallel()
	result0, err := ToAq2Preset(0)
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
	_, err = ToAq2Preset(114514)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	_, err = ToAq2Preset(-1)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	allAqs, err := GetAq2Presets(
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
	allAqs, err = GetAq2Presets(
		114514,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allAqs) != 16 {
		t.Errorf(
			"expected the length of result to be %d, got %d",
			16,
			len(allAqs),
		)
	}
	allAqs, err = GetAq2Presets(
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
