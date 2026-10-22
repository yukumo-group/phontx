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
	allBases, err := GetAllAq10Bases(
		-1,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allBases) != 0 {
		t.Errorf(
			"expected the length of all bases fetched here to be %d, got %d",
			0,
			len(allBases),
		)
	}
	allBases, err = GetAllAq10Bases(
		114514,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allBases) != 3 {
		t.Errorf(
			"expected the length of all bases fetched here to be %d, got %d",
			3,
			len(allBases),
		)
	}
	allBases, err = GetAllAq10Bases(
		2,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allBases) != 2 {
		t.Errorf(
			"expected the length of all bases fetched here to be %d, got %d",
			2,
			len(allBases),
		)
	}
}

// TestConvertIntToBase tests the convertion from integer to base
func TestConvertIntToBase(t *testing.T) {
	t.Parallel()
	expectedBases := []Aq10Base{
		F1E,
		F2E,
		M1E,
	}
	fetchedBases, err := GetAllAq10Bases(3)
	if err != nil {
		t.Error(err)
	}
	for i := range fetchedBases {
		if i == len(expectedBases) {
			t.Fatalf(
				"i became larger than length of expected bases %d",
				len(expectedBases),
			)
		}
		if fetchedBases[i] != expectedBases[i] {
			t.Errorf(
				"in idx %d, expected %s, got %s",
				i,
				expectedBases[i].ToString(),
				fetchedBases[i].ToString(),
			)
		}
	}
}
