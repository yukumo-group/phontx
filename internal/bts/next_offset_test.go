package bts

import (
	"testing"
)

// TestGetNextOffset tests getting the next offset
func TestGetNextOffset(t *testing.T) {
	t.Parallel()
	result := NextOffset(0, 201)
	if result%4 != 0 {
		t.Fatalf(
			"next offset %d cannot be devided by 4",
			result,
		)
	}
}
