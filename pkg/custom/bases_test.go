package custom

import (
	"testing"
)

// TestBases tests functions related to bases
func TestBases(t *testing.T) {
	t.Parallel()
	testData1, err := ToAq10Base(
		0,
	)
	if err != nil {
		t.Error(err)
	}
	if testData1 != F1E {
		t.Errorf(
			"expected %s, got %s",
			F1E.ToString(),
			testData1.ToString(),
		)
	}
	testData2, err := ToAq10Base(
		1,
	)
	if err != nil {
		t.Error(err)
	}
	if testData2 != F2E {
		t.Errorf(
			"expected %s, got %s",
			F2E.ToString(),
			testData2.ToString(),
		)
	}
	testData3, err := ToAq10Base(
		2,
	)
	if err != nil {
		t.Error(err)
	}
	if testData3 != M1E {
		t.Errorf(
			"expected %s, got %s",
			M1E.ToString(),
			testData3.ToString(),
		)
	}
}
