package bts

import (
	"os"
	"testing"
)

// TestBtsEquals tests the equals of bytes data
func TestBtsEquals(t *testing.T) {
	t.Parallel()
	// Define needed files
	const testFile1 string = "./testdata/aq_f1c.phont"
	const testFile2 string = "./testdata/aq_f1c_copy.phont"
	const testFile3 string = "./testdata/aq_huskey.phont"
	data1, err := os.ReadFile(
		testFile1,
	)
	if err != nil {
		t.Error(err)
	}
	data2, err := os.ReadFile(
		testFile2,
	)
	if err != nil {
		t.Error(err)
	}
	data3, err := os.ReadFile(
		testFile3,
	)
	if err != nil {
		t.Error(err)
	}
	isSame123 := IsBytesSame(
		data1,
		data2,
		data3,
	)
	if isSame123 {
		t.Error(
			"expected not to be the same",
		)
	}
	isSame12 := IsBytesSame(
		data1,
		data2,
	)
	if !isSame12 {
		t.Error(
			"file1 and file2 expected to have the same bytes",
		)
	}
	isSame13 := IsBytesSame(
		data2,
		data3,
	)
	if isSame13 {
		t.Error(
			"file1 and file3 expected not to be the same",
		)
	}
	isNilsSame := IsBytesSame(nil, nil)
	if !isNilsSame {
		t.Error(
			"expected two nils to be the same",
		)
	}
	isNilAndBytesSame := IsBytesSame(
		nil,
		data1,
	)
	if isNilAndBytesSame {
		t.Error(
			"nil and bytes should not be the same",
		)
	}
}
