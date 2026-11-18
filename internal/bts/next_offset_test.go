package bts

import (
	"testing"
)

// TestGetNextOffset tests getting the next offset
func TestGetNextOffset(t *testing.T) {
	t.Parallel()
	result1 := NextOffset(0, 201)
	if result1%4 != 0 {
		t.Fatalf(
			"next offset %d cannot be devided by 4",
			result1,
		)
	}
	if result1 < 201 {
		t.Fatalf(
			"the result1 %d fetched overlapped with %d",
			result1,
			201,
		)
	}
	result2 := NextOffset(result1, 114514)
	if result2%4 != 0 {
		t.Fatalf(
			"next offset %d cannot be devided by 4",
			result2,
		)
	}
}
