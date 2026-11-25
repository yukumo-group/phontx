package tone

import (
	"testing"

	"github.com/1Vewton/structdisplay/pkg/display"
	"github.com/yukumo-group/phontx/pkg/aqdata"
)

// TestYAMLConvertion tests the convertion of yaml data to raw tone
func TestYAMLConvertion(t *testing.T) {
	t.Parallel()
	expectedData := NewData().SetAq1Preset(
		aqdata.Aq1F1,
	).SetAq2Preset(
		aqdata.Aq2AqYukkuri,
	).SetAq10Preset(
		aqdata.Aq10Custom,
	)
	readData, err := RawYAMLFileToPhontx(
		"testdata/testdata.yaml",
	)
	if err != nil {
		t.Fatal(err)
	}
	isDataEquals := expectedData.Equals(
		readData,
	)
	if !isDataEquals {
		testDisplayer := display.NewConverter().WithRecursionLimit(
			99,
		)
		dataDisplayedExpected, err := testDisplayer.ConvertAnyData(
			"ExpectedData",
			"",
			expectedData,
			0,
		)
		if err != nil {
			t.Fatal(err)
		}
		dataDisplayedFetched, err := testDisplayer.ConvertAnyData(
			"ExpectedData",
			"",
			readData,
			0,
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Error(dataDisplayedExpected)
		t.Error(dataDisplayedFetched)
		t.Fatal(
			"fetched datas are not equal",
		)
	}
}
