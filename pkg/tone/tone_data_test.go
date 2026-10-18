package tone

import (
	"os"
	"testing"

	"github.com/yukumo-group/phontx/internal/bts"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
)

// TestToneDataSetting tests the setting of tone data
func TestToneDataSetting(t *testing.T) {
	t.Parallel()
	const testDataPath string = "./testdata/aq_f1c.phont"
	testPhont, err := os.ReadFile(
		testDataPath,
	)
	if err != nil {
		t.Error(err)
	}
	testCustomAq10Tone := custom.NewAq10CustomTone().WithFSC(
		40,
	)
	testData := NewData().SetAq1Preset(
		aqdata.Aq1F1,
	).SetAq2Preset(
		aqdata.Aq2AqTeto1,
	).SetAq10Preset(
		aqdata.Aq10F1,
	).SetAq10CustomTone(
		testCustomAq10Tone,
	)
	// Aq1Preset
	if testData.Aq1Preset == nil {
		t.Error(
			"Aq1Preset is not successfully set",
		)
	}
	aq1Preset := testData.Aq1Preset
	if *aq1Preset != aqdata.Aq1F1 {
		t.Errorf(
			"expected Aq1Preset to be %s, got %s",
			aqdata.Aq1F1.ToString(),
			(*aq1Preset).ToString(),
		)
	}
	// Aq2Preset
	if testData.Aq2Preset == nil {
		t.Error(
			"Aq2Preset is not successfully set",
		)
	}
	aq2Preset := testData.Aq2Preset
	if *aq2Preset != aqdata.Aq2AqTeto1 {
		t.Errorf(
			"expected Aq1Preset to be %s, got %s",
			aqdata.Aq2AqTeto1.ToString(),
			(*aq2Preset).ToString(),
		)
	}
	// Aq10Preset
	if testData.Aq10Preset == nil {
		t.Error(
			"Aq10Preset is not successfully set",
		)
	}
	aq10Preset := testData.Aq10Preset
	if *aq10Preset != aqdata.Aq10F1 {
		t.Errorf(
			"expected Aq1Preset to be %s, got %s",
			aqdata.Aq10F1.ToString(),
			(*aq10Preset).ToString(),
		)
	}
	// SetPhontData
	err = testData.SetPhontDataFromFile(
		testDataPath,
	)
	if err != nil {
		t.Error(err)
	}
	originalTestPhontLength := len(testPhont)
	phontFileDataLength := len(testData.PhontFileData)
	if phontFileDataLength != originalTestPhontLength {
		t.Errorf(
			"expected length %d, got %d",
			originalTestPhontLength,
			phontFileDataLength,
		)
	}
	for i := range testData.PhontFileData {
		if testData.PhontFileData[i] != testPhont[i] {
			t.Errorf(
				"expected %b on position %d, got %b",
				testPhont[i],
				i,
				testData.PhontFileData[i],
			)
		}
	}
}

// TestAllowedEngine tests the allowance of tone data to engines
func TestAllowedEngine(t *testing.T) {
	t.Parallel()
	const testDataPath string = "./testdata/aq_f1c.phont"
	testCustomAq10Tone := custom.NewAq10CustomTone().WithFSC(
		40,
	)
	// expected to support nothing
	testData1 := NewData()
	versions := aqdata.GetAllVersions()
	for _, version := range versions {
		versionAllowed, err := testData1.IsVersionAllowed(
			version,
		)
		if err != nil {
			t.Error(err)
		}
		if versionAllowed {
			t.Errorf(
				"%s is not supposed to allow",
				version.ToString(),
			)
		}
	}
	// expected to only support aq1
	testData1 = testData1.SetAq1Preset(
		aqdata.Aq1DVD,
	)
	aq1Allowed, err := testData1.IsVersionAllowed(
		aqdata.AqVer1,
	)
	if err != nil {
		t.Error(err)
	}
	if !aq1Allowed {
		t.Error(
			"expected to support aq1",
		)
	}
	// expected to only support aq2
	testData21 := NewData().SetAq2Preset(
		aqdata.Aq2AqTeto1,
	)
	isAq2Supported, err := testData21.IsVersionAllowed(
		aqdata.AqVer2,
	)
	if err != nil {
		t.Error(err)
	}
	if !isAq2Supported {
		t.Error(
			"expected to support aq2",
		)
	}
	testData22 := NewData()
	err = testData22.SetPhontDataFromFile(
		testDataPath,
	)
	if err != nil {
		t.Error(err)
	}
	isAq2Supported, err = testData22.IsVersionAllowed(
		aqdata.AqVer2,
	)
	if err != nil {
		t.Error(err)
	}
	if !isAq2Supported {
		t.Error(
			"expected to support aq2",
		)
	}
	// expected to only support aq10
	testData101 := NewData().SetAq10Preset(
		aqdata.Aq10F1,
	)
	isAq10Supported, err := testData101.IsVersionAllowed(
		aqdata.AqVer10,
	)
	if err != nil {
		t.Error(err)
	}
	if !isAq10Supported {
		t.Error(
			"expected to support aq10",
		)
	}
	testData102 := NewData().SetAq10CustomTone(
		testCustomAq10Tone,
	)
	isAq10Supported, err = testData102.IsVersionAllowed(
		aqdata.AqVer10,
	)
	if err != nil {
		t.Error(err)
	}
	if !isAq10Supported {
		t.Error(
			"expected to support aq10",
		)
	}
}

// TestDataReadLoad tests the read and load of data
func TestDataReadLoad(t *testing.T) {
	t.Parallel()
	// Prepare data
	const testDataPath string = "./testdata/aq_f1c.phont"
	tmpDir := t.TempDir()
	fileName := "test"
	testCustomAq10Tone := custom.NewAq10CustomTone().WithFSC(
		40,
	).WithBase(
		custom.M1E,
	)
	newData := NewData().SetAq1Preset(
		aqdata.Aq1DVD,
	).SetAq2Preset(
		aqdata.Aq2AqDefo1,
	).SetAq10Preset(
		aqdata.Aq10Custom,
	).SetAq10CustomTone(
		testCustomAq10Tone,
	)
	err := newData.SetPhontDataFromFile(
		testDataPath,
	)
	if err != nil {
		t.Error(err)
	}
	// Load and save
	filePath, err := newData.SaveFile(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	reloadedData, err := NewDataFromFile(
		filePath,
	)
	if err != nil {
		t.Error(err)
	}
	if !reloadedData.Equals(newData) {
		t.Error(
			"expected the loaded and new data to be the same",
		)
		if !bts.IsBytesSame(
			reloadedData.PhontFileData,
			newData.PhontFileData,
		) {
			t.Error(
				"the two tones do not have the same phont data read",
			)
		}
		if reloadedData.Aq1Preset != newData.Aq1Preset {
			t.Errorf(
				"expected %s, got %s",
				newData.Aq1Preset.ToString(),
				reloadedData.Aq1Preset.ToString(),
			)
		}
	}
}
