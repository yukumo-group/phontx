package tone

import (
	"testing"

	"github.com/yukumo-group/phontx/pkg/aqdata"
)

// TestToneDataSetting tests the setting of tone data
func TestToneDataSetting(t *testing.T) {
	t.Parallel()
	testData := NewData().SetAq1Preset(
		aqdata.Aq1F1,
	).SetAq2Preset(
		aqdata.Aq2AqTeto1,
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
}
