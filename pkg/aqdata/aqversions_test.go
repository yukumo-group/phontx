package aqdata

import (
	"testing"
)

// TestConvertIntToAqVersion tests the converting from int to aquestalk version
func TestConvertIntToAqVersion(t *testing.T) {
	t.Parallel()
	result1 := ToAquesVersion(
		1,
	)
	if result1 != AqVer1 {
		t.Errorf(
			"expected %s, got %s",
			AqVer1.ToString(),
			result1.ToString(),
		)
	}
	result2 := ToAquesVersion(
		2,
	)
	if result2 != AqVer2 {
		t.Errorf(
			"expected %s, got %s",
			AqVer2.ToString(),
			result2.ToString(),
		)
	}
	result10 := ToAquesVersion(
		10,
	)
	if result10 != AqVer10 {
		t.Errorf(
			"expected %s, got %s",
			AqVer10.ToString(),
			result10.ToString(),
		)
	}
	resultInt1 := AqVer1.ToInt()
	if resultInt1 != 1 {
		t.Errorf(
			"expected %d, got %d",
			1,
			resultInt1,
		)
	}
	resultInt2 := AqVer2.ToInt()
	if resultInt2 != 2 {
		t.Errorf(
			"expected %d, got %d",
			2,
			resultInt2,
		)
	}
	resultInt10 := AqVer10.ToInt()
	if resultInt10 != 10 {
		t.Errorf(
			"expected %d, got %d",
			10,
			resultInt10,
		)
	}
}
